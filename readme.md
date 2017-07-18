requirements:
- `brew install kubernetes-cli`
- `brew cask install minikube`

start minikube: `minikube start`

sets up docker env variables: `eval $(minikube docker-env --use-vendored-driver)`

build docker image: `docker build -t example:1.0 .`

deploy deployment: `kubectl create -f deployment.yml`

deploy service: `kubectl create -f service.yml`

---
in browser, go to `http://192.168.99.100:31000`


deploy new image to deployment
`kubectl set image deployment/example example=example:1.5`
