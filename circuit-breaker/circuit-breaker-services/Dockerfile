FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app main.go

CMD ["/app/app"]
