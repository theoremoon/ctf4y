FROM golang:1.13-alpine

ADD challenge/templates /templates

RUN mkdir /build
WORKDIR /build
ADD challenge/main.go /build/main.go
ADD challenge/go.mod /build/go.mod
ADD challenge/go.sum /build/go.sum

RUN go mod download
RUN go build -o /app
WORKDIR /

CMD /app
