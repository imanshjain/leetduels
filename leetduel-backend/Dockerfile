# Build stage
FROM golang:1.24.4-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from build stage
COPY --from=build /app/main .
# Copy any necessary configuration files or assets
COPY --from=build /app/SECRETS /app/SECRETS

# Expose port for the application
EXPOSE 8080

# Run the binary
CMD ["./main"]