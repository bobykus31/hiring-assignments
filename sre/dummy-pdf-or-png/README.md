# dummy-pdf-or-png microservice
Pseudorandomly returns either a pdf or png file.

Responds to GET requests on "/" port 3040.

Dockerfile multi-stage takes care of building, so to get started you run the following commands in this directory, assuming you have Docker installed.

```sh
docker build -t dummy-pdf-or-png .
docker run --rm -it -p 3030:3030 dummy-pdf-or-png
```

On purpose, the service will sometimes deliver a corrupt pdf file.

CT deploy 
```
docker build -t dummy-pdf-or-png .
docker tag dummy-pdf-or-png bobykus/dummy-pdf-or-png:1.0.1
docker push bobykus/dummy-pdf-or-png:1.0.1
```

Minikube part

```
minikube start
kubectl apply -f k8s-sre-deployment.ym
kubectl get pods
kubectl get services
minikube service  dummy-pdf-or-png-service --url
```

Logs can be shipped to  timber
```
https://docs.timber.io/setup/platforms/kubernetes#automatic-context

   kubectl create namespace logging
   kubectl create -f https://raw.githubusercontent.com/fluent/fluent-bit-kubernetes-logging/master/fluent-bit-service-account.yaml
   kubectl create -f https://raw.githubusercontent.com/fluent/fluent-bit-kubernetes-logging/master/fluent-bit-role.yaml
   kubectl create -f https://raw.githubusercontent.com/fluent/fluent-bit-kubernetes-logging/master/fluent-bit-role-binding.yaml

   kubectl create -f https://gist.githubusercontent.com/binarylogic/951ea32ed462933fa70c439f9cab06f3/raw/fe6b761770f1ccd9b44c96421d4892b8e96927d6/fluent-bit-configmap.yaml

   kubectl create -f  fluent-bit-ds.yaml
```
