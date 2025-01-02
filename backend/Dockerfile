FROM golang:1.23-alpine
WORKDIR /app

# Copy .env file first
COPY .env .env

# Copy the rest of the application
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
EXPOSE 8000
CMD ["./main"]
