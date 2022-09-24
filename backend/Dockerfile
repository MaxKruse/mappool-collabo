FROM golang:1.18

WORKDIR /app

# Install gcc as a sqlite3 dependency
RUN apt-get update && apt-get install -y gcc

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 5000

CMD ["/app/main"]
