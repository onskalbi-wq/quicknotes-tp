FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY app/go.mod ./
COPY app/main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/quicknotes .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/quicknotes .
EXPOSE 8080
CMD ["./quicknotes"]
