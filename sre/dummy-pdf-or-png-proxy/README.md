# dummy-pdf-or-png-proxy microservice

Takes HTTP GET requests with a random ID (/1, /529852, etc.), requests a document from the microservice  dummy-pdf-or-png, and then returns the document with correct mime-type.
Microservice  dummy-pdf-or-png supposed to be running and available on some URL. This URL is passed to dummy-pdf-or-png-proxy as http_proxy env

Requests can be run as

```
http_proxy=http://dummy-pdf-or-png:3030 curl  http://dummy-pdf-or-png-proxy:3444
```

or whatever port is exposed  by k8s.
