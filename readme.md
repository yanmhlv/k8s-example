requirements:
- `brew install kubernetes-cli`
- `brew cask install minikube`

start minikube: `minikube start`

use local docker daemon: `eval $(minikube docker-env --use-vendored-driver)`

---
`cat main.go`
```
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world %d", n)
	})

	(&http.Server{
		Addr:    ":3000",
		Handler: mux,
	}).ListenAndServe()
}
```

----
`cat Dockerfile`
```
FROM golang:1.8.3-alpine3.6

COPY . /go/src/app
RUN go install app
EXPOSE 3000
CMD ["/go/bin/app"]
```
build docker image: `docker build -t example:1.0 .`

---

`cat deployment.yml`
```
kind: Deployment
apiVersion: apps/v1beta1
metadata:
  name: example
spec:
  replicas: 10
  template:
    metadata:
      labels:
        app: example
    spec:
      containers:
        - name: example
          imagePullPolicy: Never
          image: example:1.2
          ports:
            - name: http
              containerPort: 3000

```
deploy deployment: `kubectl create -f deployment.yml`

---

`cat service.yml`
```
kind: Service
apiVersion: v1
metadata:
  name: example
spec:
  selector:
    app: example
  type: NodePort
  ports:
    - port: 3000
      nodePort: 31000
```
deploy service: `kubectl create -f service.yml`

---
in browser, go to `http://192.168.99.100:31000`
