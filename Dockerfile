# Build stage
FROM golang:1.22-alpine as build

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

RUN go build -ldflags="-linkmode external -extldflags -static" -tags netgo -o main cmd/server/main.go

# Final stage
FROM scratch

ENV DB_TYPE=sqlite
ENV DB_STORAGE=memory
ENV GIN_MODE=test
ENV TRUSTED_PROXIES=127.0.0.1

COPY --from=build /app/main /main

EXPOSE 8080

CMD ["/main"]
