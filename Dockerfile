# Build the Go application
FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Run the Go application in slim image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]