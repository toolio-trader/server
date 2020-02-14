# server

Simple tooling to expose a folder as www root for serving web content 

use binary : 

## HTTP server

./server --port 8000 --path www

Exposes folder www via port 8000


## A+ HTTPS Server, ssllabs

./serverSSL --domains 'dashboard.toolio.dev' --path www

Exposes folder www via port 443 with ssl encryption

Make sure user has read/write access in /tmp/cache-golang-autocert-... folder!

-- A+ notification on https://www.ssllabs.com/ssltest/

ADD DNS RECORDS:
CAA RECORD: 128 issue "letsencrypt.org"


or build your own, golang must be installed on host system for compilation:

HTTP Server  : go build server.go
HTTPS Server : go build serverSSL.go
