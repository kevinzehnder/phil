
## BUILD
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main ./cmd/main/main.go

## RUN
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin/main /app/main
WORKDIR /app
EXPOSE 9000

ENTRYPOINT ["/app/main"]
