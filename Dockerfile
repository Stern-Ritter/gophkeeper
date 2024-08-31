FROM golang:1.22 AS builder

WORKDIR /app/gophkeeper

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /gophkeeper ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /gophkeeper /app/gophkeeper

COPY --from=builder /app/gophkeeper/certs /app/certs

RUN mkdir -p /app/files

CMD ["/app/gophkeeper"]

