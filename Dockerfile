FROM golang:1.22-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY app /app

ENV CGO_ENABLED=1

RUN go build -o main cmd/server/main.go

EXPOSE 8080

CMD ["./main"]
