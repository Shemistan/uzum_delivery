FROM golang:1.21.1-alpine AS builder

COPY . /app/
WORKDIR /app

RUN go mod download
RUN go build -o /bin/uzum_delivery_service ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /bin/uzum_delivery_service .

CMD ["./uzum_delivery_service"]
