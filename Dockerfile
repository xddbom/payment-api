FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN cd cmd/payment-api && CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/cmd/payment-api/main /main

EXPOSE 8080
CMD ["/main"]