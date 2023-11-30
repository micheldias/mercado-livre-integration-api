FROM golang:1.21-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata openssh make
WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -installsuffix cgo -o ./main cmd/server/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
WORKDIR /app
COPY --from=builder /app/application.yaml application.yaml
COPY --from=builder /app/main ./main

ENTRYPOINT ["/app/main"]