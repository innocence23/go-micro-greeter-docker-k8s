
eval $(minikube docker-env)

docker build -t greeter-service ./greeter-service

docker build -t greeter-client ./greeter-client



# ❯ kubectl run -it --rm --image=curlimages/curl --restart=Never one-time-curl -- http://192.168.49.2:30012/greeter\?name\=1abc
# minikube start --memory=4096 --driver=virtualbox


# ❯ minikube service greeter-client --url -n go-micro  可用http://127.0.0.1:54412/greeter?name=1abc


# kubectl port-forward svc/greeter-client 3000:80 -n go-micro   http://127.0.0.1:3000/greeter?name=1abc