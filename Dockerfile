FROM golang:1.8.3

WORKDIR /app
COPY . /app

ENV GOPATH=/app:/app/vendor
RUN cd /app \
    && go get service.1 \
    && go get service.2 \
    && go get service.3

EXPOSE 3000
