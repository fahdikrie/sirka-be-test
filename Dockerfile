FROM golang:1.18
WORKDIR /go/src/sirka-be-test
COPY . .
RUN go build -o bin/server src/main.go
CMD ["./bin/server"]
