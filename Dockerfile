FROM golang:1.22

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install golang-migrate for database migrations
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy application source code, Swaggo docs, config, and migrations
COPY . ./

# Build the Go application (assuming main.go is in cmd/web directory)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/chat-app-service ./cmd/web/main.go

# Expose the application port
EXPOSE 9000

# Define the MYSQL_URL environment variable (update 'localhost' if using Docker Compose)
ENV MYSQL_URL=mysql://root:password@tcp(database:3306)/chat-app?charset=utf8mb4&parseTime=True&loc=Local

# Run migrations before starting the app
CMD migrate -database "${MYSQL_URL}" -path /app/db/migrations up && /app/chat-app-service
