# Build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-docs .

# Runtime image
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache \
    curl \
    docker-cli \
    git \
    bash \
    ca-certificates

# Install syft
RUN curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /usr/local/bin

# Install grype
RUN curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /usr/local/bin

# Install dive
RUN curl -L https://github.com/wagoodman/dive/releases/download/v0.12.0/dive_0.12.0_linux_amd64.tar.gz -o dive.tar.gz && \
    tar -xzf dive.tar.gz -C /usr/local/bin dive && \
    rm dive.tar.gz

COPY --from=builder /app/docker-docs /usr/local/bin/docker-docs

ENTRYPOINT ["docker-docs"]
