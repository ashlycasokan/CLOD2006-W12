# Insecure: Using a large, vulnerable base image
FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o app .
CMD ["./app"]
