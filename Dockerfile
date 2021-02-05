FROM golang:1.15.8

VOLUME /ginkgo-experiments
WORKDIR /ginkgo-experiments

ADD go.mod .
ADD go.sum .

RUN go mod download

# Install ginkgo command
RUN go get github.com/onsi/ginkgo/ginkgo
