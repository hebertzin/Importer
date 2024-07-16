FROM golang:1.22.0-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o enube-challenge ./cmd/main.go

EXPOSE 8080

CMD ["./enube-challenge"]