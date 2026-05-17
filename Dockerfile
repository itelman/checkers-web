# Build Stage
FROM golang:1.26-alpine AS builder
RUN apk add --no-cache build-base gcc musl-dev

WORKDIR /app

COPY . .
RUN go mod download

# Final Stage (no need to copy binary)
FROM golang:1.26-alpine
WORKDIR /app

# Copy application code from the builder stage
COPY --from=builder /app .

EXPOSE 8888

# Start the application using go run
CMD ["go", "run", "./api"]
