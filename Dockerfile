FROM golang:1.8.3-alpine3.6 AS build-env

COPY . /go/src/app
RUN go install app

FROM alpine
COPY --from=build-env /go /go
EXPOSE 3000
CMD ["/go/bin/app"]
