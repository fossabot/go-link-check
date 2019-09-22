#-------------------------------------------------------------------------------
#
# Docker Configuration, https://docs.docker.com/engine/reference/builder/
#
# Includes "builder" stage to compile and test application and slimmed down
# stage to run the application.
#
#-------------------------------------------------------------------------------

FROM golang:1.13.0-alpine3.10 as builder

LABEL intermediate=true

RUN apk add --no-cache gcc

COPY . /app
WORKDIR /app

# Install module dependencies, run test suite, and build executable.
RUN go mod vendor \
  && go build -v -mod=vendor -o ./go-link-check

FROM alpine:3.10.2

ARG BUILD_DATE=""
ARG VCS_REF=""
ARG VERSION=""

# https://github.com/opencontainers/image-spec/blob/master/annotations.md
LABEL \
  org.opencontainers.image.authors="Daniel Tedman" \
  org.opencontainers.image.created="${BUILD_DATE}" \
  org.opencontainers.image.description="Link check automation tool." \
  org.opencontainers.image.documentation="https://github.com/dbtedman/go-link-check" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.revision="${VCS_REF}" \
  org.opencontainers.image.source="https://github.com/dbtedman/go-link-check/commit/${VCS_REF}" \
  org.opencontainers.image.title="Go Link Check" \
  org.opencontainers.image.url="https://github.com/dbtedman/go-link-check" \
  org.opencontainers.image.vendor="Daniel Tedman" \
  org.opencontainers.image.version="${VERSION}"

WORKDIR /app

COPY --from=builder /app /app/

# Configure docker to execute our application on run.
ENTRYPOINT ["./go-link-check"]
