FROM golang:1.18.0-alpine AS builder

WORKDIR /app
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./docs ./docs
COPY ./error ./error
COPY ./logger ./logger
COPY ./v1 ./v1
COPY ./vendor ./vendor
COPY ./echo.go ./echo.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go
COPY ./router.go ./router.go

RUN apk --no-cache add ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main" -ldflags="-w -s" ./cmd/main.go

FROM scratch

COPY --from=builder /app/main /usr/bin/
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

CMD ["main"]

EXPOSE 80