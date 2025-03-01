FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY /migrations /app/migrations
WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
FROM alpine:latest
COPY --from=builder /app/cmd/server/server /server
COPY --from=builder /app/migrations /migrations
EXPOSE 8080

ENTRYPOINT ["/server"]
