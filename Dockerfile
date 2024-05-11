# For build
FROM golang:1.21.5-alpine3.19 AS builder

RUN apk update \
    && apk add --no-cache git gcc g++ \
    && apk upgrade \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main ./cmd/anytore-backend/


# For release
FROM alpine:3.19
COPY --from=builder /app/main /app/main

ENV LOG_LEVEL=info \
    GIN_MODE=release \
    PORT=80 \
    AWS_REGION=ap-northeast-1

EXPOSE 80
CMD ["/app/main"]
