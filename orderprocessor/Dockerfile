FROM golang:latest AS build

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
RUN go mod download

WORKDIR ./cmd

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /orderprocessor

# Final stage
FROM alpine:latest

COPY --from=build /orderprocessor /app/
WORKDIR /app

# Run
CMD ["./orderprocessor"]

