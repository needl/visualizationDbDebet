# syntax=docker/dockerfile:1

FROM golang:1.26.2-alpine3.23 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/api

FROM alpine:3.23.4 AS runtime

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app

COPY --from=builder /out/app /app/app

ENV DB_HOST=""
ENV DB_PORT=5432
ENV DB_USER=""
ENV DB_PASSWORD=""
ENV DB_NAME=""
ENV DB_SSLMODE=disable
ENV PORT=8181

EXPOSE 8181

USER app

ENTRYPOINT ["/app/app"]
