FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app cmd/main.go

CMD ["./app"]