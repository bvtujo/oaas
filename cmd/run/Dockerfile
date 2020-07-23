FROM golang:1.13
WORKDIR /otter
COPY . .
ENV GOPROXY=direct
RUN go build otter.go
ENTRYPOINT go run otter.go
EXPOSE 8080