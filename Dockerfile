FROM golang:alpine3.18 AS builder
COPY goweb.go src/
RUN go build -o goweb src/goweb.go

FROM alpine:3.18
LABEL org.opencontainers.image.authors="carmelo[DOT]califano[AT]gmail[DOT]com"
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/goweb ./
CMD ["./goweb"]

