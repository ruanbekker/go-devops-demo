FROM golang:1.22-alpine

ENV DB_TYPE=sqlite
ENV DB_STORAGE=memory
ENV GIN_MODE=test
ENV TRUSTED_PROXIES=127.0.0.1

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY app /app

ENV CGO_ENABLED=1

RUN go build -o main cmd/server/main.go

EXPOSE 8080

CMD ["./main"]
