FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build optimized binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o air_trail_backend .

# Stage 2: minimal image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/air_trail_backend .
COPY --from=builder /app/static ./static
COPY .env .

# non-root user
RUN adduser -D appuser
USER appuser

# Load env + run app
CMD ["./air_trail_backend"]
