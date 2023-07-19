package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func parseIP(peer string) string {
	addr := peer[:strings.LastIndex(peer, ":")]

	if strings.Contains(addr, "[") {
		return addr[1 : len(addr)-1]
	}
	return addr
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	RemoteAddr := parseIP(r.RemoteAddr)
	log.Printf("GET / from %s\n", r.RemoteAddr)
	w.Header().Add("X-ADDR", RemoteAddr)
	io.WriteString(w, RemoteAddr)
}

func main() {

	http.HandleFunc("/", getRoot)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "3333"
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Printf("Server is listening *%s", fmt.Sprintf(":%s", port))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server closed")

}
