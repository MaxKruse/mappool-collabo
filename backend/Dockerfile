# builder
FROM golang:1.18-alpine AS builder

WORKDIR /app

# Install gcc as a sqlite3 dependency
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 5000

CMD ["/app/main"]
