FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .
CMD ["./main"]
EXPOSE 8082
