FROM golang:1.23.4

WORKDIR /payment-api

COPY . .

RUN go build -o main .

CMD ["./main"]
