# server

use binary : 

Simple server

./server --port 8000 --path www

Exposes folder www via port 8000

or build your own, golang must be installed on host system for compilation:

go build server.go


Simple server SSL A+ ssllabs

./serverSSL --domains 'dashboard.toolio.dev' --path www

Exposes folder www via port 443 with ssl encryption

ADD DNS RECORDS for A+ :

CAA RECORD: 128 issue "letsencrypt.org"


or build your own, golang must be installed on host system for compilation:

go build serverSSL.go
