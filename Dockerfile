FROM golang:1.26-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /usr/src/app/main .
COPY --from=builder /usr/src/app/web ./web

EXPOSE 8080
CMD ["./main"]

