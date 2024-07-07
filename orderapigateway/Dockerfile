FROM golang:latest AS build

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN go mod download

WORKDIR ./cmd

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /orderapigateway

# Final stage
FROM alpine:latest

COPY --from=build /orderapigateway /app/
WORKDIR /app

# Run
CMD ["./orderapigateway"]

