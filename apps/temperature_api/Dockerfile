FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app/app
WORKDIR /app
CMD ["./app"]
