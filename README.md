# My External Address
It shows external IP address either v4 or v6.
You can check http://4.ltecode.com or http://6.ltecode.com 
Additionally it adds new header in the response `X-ADDR` with value of the IP.

## Install the service

Clone the repository

```bash
cd myaddr
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o myaddrsrv
mkdir -p /opt/myaddr
mv myaddrsrv /opt/myaddr/
chmod +x /opt/myaddr/myaddrsrv
cp myaddr.service /opt/myaddr/

ln -s /opt/myaddr/myaddr.service  /etc/systemd/system/myaddr.service 
systemctl daemon-reload
systemctl enable myaddr.service
systemctl is-enabled myaddr.service
systemctl start myaddr.service
systemctl status myaddr.service
```

## Environemnt 
In the service file there is  `Environment="HTTP_PORT=80"` it defines the port to listen incoming request.

or to run it without systemctl

```
export HTTP_PORT=80
/opt/myaddr/myaddrsrv
```

# Client

In the `client` folder you can find he.py script which will help you update IPv4 tunnel broker endpoint.
You have to point your `login`, `update key` and `tunnel id` you can take them from the https://www.tunnelbroker.net site on the Advanced tab.

You can use the client on the remote side where the IPv4 is dynamic.

## Installation of client

```
cp client/he.py /opt/myaddr/he.py
cp client/he.service /opt/myaddr/he.service
ln -s /opt/myaddr/he.service  /etc/systemd/system/he.service 
systemctl daemon-reload
systemctl enable he.service
systemctl is-enabled he.service
systemctl start he.service
systemctl status he.service
```