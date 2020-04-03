FROM golang:1.13
WORKDIR /otter
COPY . .
RUN go build otter.go
ENTRYPOINT go run otter.go
EXPOSE 8080
