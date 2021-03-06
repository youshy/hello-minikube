# Hello Minikube

AKA Small intro to Kubernetes (now for real)

## Run minikube

* `minikube start`
* `minikube dashboard`

## Run with local image

* `eval $(minikube docker-env)`
* Remember to tag your build `docker build -t hello-minikube:latest`
* Run with `imagePullPolicy` set to `Never` - `kubectl run hello-minikube –-image=hello-minikube:latest -–image-pull-policy=Never`

## Commands

* View deployment - `kubectl get deployments`
* View pod - `kubectl get pods`
* View cluster events - `kubectl get events`
* Delete deployment - `kubectl delete -n default deployment hello-minikube`
* Expose port - `kubectl expose deployment hello-minikube --type=LoadBalancer --port=8080`
* View services - `kubectl get services`
* (If on minikube) - `minikube service hello-minikube` (opens up the service in the browser)

### Cleanup

* `kubectl delete service hello-minikube`
* `kubectl delete deployment hello-minikube`
* `minikube stop`
