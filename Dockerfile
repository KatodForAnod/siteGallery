FROM golang:latest
WORKDIR /app
COPY . .
RUN go build cmd/main.go
CMD ["/app/main"]