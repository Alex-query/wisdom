FROM golang
USER root

ENV GO111MODULE on

WORKDIR /go/src/service
COPY . .

RUN go mod vendor
RUN go build -mod vendor -o ./service .

ENTRYPOINT ./service server