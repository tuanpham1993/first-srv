FROM golang:alpine
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo main.go
ENTRYPOINT [ "/app/main" ]
