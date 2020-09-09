FROM golang:1.14

WORKDIR /go/src/app
COPY openshift-test.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]