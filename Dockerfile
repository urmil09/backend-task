# Build stage
FROM golang:1.22-alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy all files into the container
COPY . /app

# Install dependencies
RUN go mod tidy && go mod vendor

# Build the Go binary
RUN go build -o flink-backend-assignment ./cmd/api

# Final runtime stage
FROM alpine:latest

# Set environment variables previously in docker-compose
ENV POSTGRES_USER=flink
ENV POSTGRES_PASSWORD=flink.8080
ENV POSTGRES_DB=flink
ENV DB_HOST=postgres-db
ENV HTTP_PORT=8080

# Expose the HTTP port
EXPOSE $HTTP_PORT

# Copy the binary from the builder stage
COPY --from=builder /app/flink-backend-assignment .

# Set the entrypoint command
CMD ["./flink-backend-assignment"]
