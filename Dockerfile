FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 4040

CMD ["./main"]