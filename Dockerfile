# Base image
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

#build binary
RUN go build -o main .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"]