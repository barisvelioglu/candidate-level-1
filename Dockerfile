FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o sensor-simulator

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sensor-simulator .

EXPOSE 8080
CMD ["./sensor-simulator"] 