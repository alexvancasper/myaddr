#!/usr/bin/python3

import time
import requests
import base64

LOGIN = ""
UPDATE_KEY = ""
TUNNEL_HOST = ""
HEurl = 'https://ipv4.tunnelbroker.net/nic/update?hostname={0}'.fromat(TUNNEL_HOST)
IDENT = 'http://4.ltecode.com'
SLEEP_TIME = 3 # minutes

last_v4ip = ""

b = base64.b64encode(bytes('{0}:{1}'.format(LOGIN, UPDATE_KEY), 'utf-8'))
base64_str = b.decode('utf-8')

while True:
    try:
       response = requests.get(IDENT)
    except Exception as e:
         print("URL {0} cannot be open error {1} ".format(IDENT, e))
         exit(1)

    if response.status_code != 200:
       print("Something wrong with URL {0} code: {1}\n".format(IDENT, e))
       exit(1)

    v4_external_ip = response.text

    if v4_external_ip != last_v4ip:
       try:
          result =  requests.get(HEurl, headers={'Authorization': 'Basic {0}'.format(base64_str)})
       except Exception as e:
          print("URL {0} cannot be open error {1} ".format(HEurl, e))
          exit(1)
       if result.status_code != 200:
          print("status-code {0}\n".format(result.status_code))
          exit(1)
       last_v4ip = v4_external_ip

       print(result.text)
    else:
       print("equal\n")

    time.sleep(SLEEP_TIME * 60)
