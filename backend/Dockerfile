FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o backend -tags prod cmd/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/backend .

CMD ["./backend"]
