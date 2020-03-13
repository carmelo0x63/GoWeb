FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main.go .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main.go ./
CMD ["./main.go"]
