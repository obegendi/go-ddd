FROM golang:1.18 as builder

RUN mkdir -p /go/src/github.com/obegendi/go-ddd
ADD . /go/src/github.com/obegendi/go-ddd
RUN ls -la
ENV GOPATH /go
ENV GO111MODULE=on
WORKDIR /go/src/github.com/obegendi/go-ddd
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN  go vet && CGO_ENABLED=0 go build -o main main.go


FROM alpine:3.11
RUN mkdir -p /app
COPY --from=builder /go/src/github.com/obegendi/go-ddd/config.local.toml /app/
COPY --from=builder /go/src/github.com/obegendi/go-ddd/main /app/

WORKDIR /app
EXPOSE 8080
ENTRYPOINT /app/main  