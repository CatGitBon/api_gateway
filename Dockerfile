FROM golang:1.24.4-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o api_gateway ./cmd

EXPOSE 50051

CMD ["./api_gateway"]