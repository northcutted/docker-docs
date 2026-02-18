# Build stage — downloads external tools only (no Go compilation needed;
# GoReleaser builds the binary separately).
FROM alpine:3.21 AS builder

WORKDIR /tmp/tools

# Install download utilities
RUN apk add --no-cache curl tar

# Pin tool versions for reproducibility.
ARG SYFT_VERSION=1.42.1
ARG GRYPE_VERSION=0.108.0
ARG DIVE_VERSION=0.13.1
ARG DOCKER_VERSION=29.2.1
ARG TARGETARCH

# Install syft (pinned version instead of latest-from-script)
RUN curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /tmp/tools v${SYFT_VERSION}

# Install grype (pinned version instead of latest-from-script)
RUN curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /tmp/tools v${GRYPE_VERSION}

# Install dive (architecture-specific)
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        curl -L https://github.com/wagoodman/dive/releases/download/v${DIVE_VERSION}/dive_${DIVE_VERSION}_linux_arm64.tar.gz -o dive.tar.gz; \
    else \
        curl -L https://github.com/wagoodman/dive/releases/download/v${DIVE_VERSION}/dive_${DIVE_VERSION}_linux_amd64.tar.gz -o dive.tar.gz; \
    fi && \
    tar -xzf dive.tar.gz dive && \
    rm dive.tar.gz

# Get Docker CLI (static binary)
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        curl -fsSLO https://download.docker.com/linux/static/stable/aarch64/docker-${DOCKER_VERSION}.tgz; \
    else \
        curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz; \
    fi && \
    tar xzvf docker-${DOCKER_VERSION}.tgz --strip 1 -C /tmp/tools docker/docker && \
    rm docker-${DOCKER_VERSION}.tgz

# Final Stage: Distroless
FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /

# Copy our Go binary
# GoReleaser v2 with buildx places artifacts in platform-specific folders
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/dock-docs /usr/local/bin/dock-docs

# Copy external tools
COPY --from=builder /tmp/tools/syft /usr/local/bin/syft
COPY --from=builder /tmp/tools/grype /usr/local/bin/grype
COPY --from=builder /tmp/tools/dive /usr/local/bin/dive
COPY --from=builder /tmp/tools/docker /usr/local/bin/docker

# Add /usr/local/bin to PATH
ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["dock-docs"]
