FROM golang:1.13
WORKDIR /otter
COPY . .
RUN go run main.go