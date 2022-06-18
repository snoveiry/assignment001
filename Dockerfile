FROM golang:1.18.0-alpine AS builder

WORKDIR /app
COPY ./assets ./assets
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./database ./database
COPY ./docs ./docs
COPY ./email ./email
COPY ./error ./error
COPY ./geo ./geo
COPY ./logger ./logger
COPY ./middleware ./middleware
COPY ./sms ./sms
COPY ./testing ./testing
COPY ./util ./util
COPY ./v2 ./v2
COPY ./v3 ./v3
COPY ./vendor ./vendor
COPY ./echo.go ./echo.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go
COPY ./router.go ./router.go

RUN apk --no-cache add ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main" -ldflags="-w -s" ./cmd/monolith/main.go

FROM scratch

COPY --from=builder /app/main /usr/bin/
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

CMD ["main"]

EXPOSE 80