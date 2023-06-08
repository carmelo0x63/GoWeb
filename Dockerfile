FROM golang:alpine3.18 AS builder
COPY main.go src/
RUN go build -o main src/main.go

FROM alpine:3.18
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/main ./
CMD ["./main"]

