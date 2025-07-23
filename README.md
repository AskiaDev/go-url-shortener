# Go URL Shortener

A simple URL shortener service built with Go, Gin, and Redis.

## Running with Docker Compose

### Prerequisites

-   Docker
-   Docker Compose

### Quick Start

1. **Build and run the services:**

    ```bash
    docker-compose up --build
    ```

2. **Run in background (detached mode):**

    ```bash
    docker-compose up -d --build
    ```

3. **Stop the services:**

    ```bash
    docker-compose down
    ```

4. **View logs:**
    ```bash
    docker-compose logs app
    docker-compose logs redis
    ```

### API Endpoints

Once running, the service will be available at `http://localhost:8080`

-   **GET /** - Welcome message
-   **POST /create-short-url** - Create a short URL
-   **GET /:shortUrl** - Redirect to original URL

### Example Usage

Create a short URL:

```bash
curl -X POST http://localhost:8080/create-short-url \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

### Development

For local development without Docker:

1. Start Redis locally on port 6379
2. Run: `go run main.go`

The application will automatically use localhost:6379 when no environment variables are set.
