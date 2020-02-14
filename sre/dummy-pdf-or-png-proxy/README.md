# dummy-pdf-or-png-proxy microservice


Takes HTTP GET requests with a random ID (/1, /529852, etc.), requests a document from the microservice we have provided in the dummy-pdf-or-png subdirectory of this repository, and then returns the document with correct mime-type.

Provides an endpoint for health monitoring /health

Has a Kubernetes Manifest dummy-pdf-or-png-proxy-deployment.yml

Has Readiness and LivenessChecks.

Has a Dockerfile.

The service log relevant information
```
kubectl logs --namespace=dummy-pdf-or-png dummy-pdf-or-png-proxy-6664d4cbb-hd6dw
2020/02/14 11:15:47 172.17.0.1:21307   GET   http://dummy-pdf-or-png-service:3040/
2020/02/14 11:15:47 172.17.0.1:21307   200 OK
2020/02/14 11:15:47 map[Accept-Ranges:[bytes] Content-Length:[13264] Content-Type:[application/pdf] Date:[Fri, 14 Feb 2020 11:15:47 GMT] Last-Modified:[Mon, 10 Feb 2020 15:40:51 GMT]]
```

Prometheus  metrics  available on /metrics URL
```
# HELP http_requests_total Count of all HTTP requests
# TYPE http_requests_total counter
http_requests_total{code="200",method="get"} 2
# HELP version Version information about this binary
# TYPE version gauge
version{version="v0.1.0"} 0
```

