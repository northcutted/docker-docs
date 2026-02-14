# syntax=docker/dockerfile:1
FROM alpine:3.19

# -----------------------------------------------------------------------------
# BUILD ARGUMENTS
# -----------------------------------------------------------------------------

# @name: VERSION
# @description: The application version to build.
# @default: 1.0.0
ARG VERSION=1.0.0

# @name: BUILD_DATE
# @description: The date the image was built (RFC3339).
# @required: true
ARG BUILD_DATE

# -----------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# -----------------------------------------------------------------------------

# @name: APP_ENV
# @description: The environment the app is running in (dev, staging, prod).
# @default: production
ENV APP_ENV="production"

# @name: DATABASE_URL
# @description: Connection string for the primary database.
# @required: true
ENV DATABASE_URL=""

# @name: LOG_LEVEL
# @description: Global logging level. Options: DEBUG, INFO, WARN, ERROR.
# @default: INFO
ENV LOG_LEVEL INFO

# Testing the parser's ability to handle equality vs space separation
# @name: API_TIMEOUT
# @description: Timeout in seconds for external API calls.
ENV API_TIMEOUT=30

# -----------------------------------------------------------------------------
# COMPLEX CASES (The "Torture" Section)
# -----------------------------------------------------------------------------

# @name: FEATURE_FLAGS
# @description: A JSON string of enabled feature flags.
# @default: {"new_ui": false}
ENV FEATURE_FLAGS='{"new_ui": false, "beta_access": true}'

# @name: PATH_additions
# @description: Extensions to the system path.
# This is a multi-line instruction test.
ENV PATH="/opt/myapp/bin:$PATH" \
    LD_LIBRARY_PATH="/opt/myapp/lib"

# @name: UNDOCUMENTED_VAR
# This variable has a comment but no magic @description, should likely be ignored or parsed as empty desc.
ENV SECRET_SAUCE="szechuan"

# This variable has NO comments at all.
ENV RAW_VAR="raw"

# -----------------------------------------------------------------------------
# METADATA & PORTS
# -----------------------------------------------------------------------------

# @name: org.opencontainers.image.authors
LABEL org.opencontainers.image.authors="platform-team@example.com"

# @name: HTTP Service
# @description: The main web server port.
EXPOSE 8080

# @name: Metrics
# @description: Prometheus metrics endpoint.
EXPOSE 9090/tcp

CMD ["./myapp"]