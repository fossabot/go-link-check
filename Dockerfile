#-------------------------------------------------------------------------------
#
# Docker Configuration, https://docs.docker.com/engine/reference/builder/
#
# Includes "builder" stage to compile and test application and slimmed down
# stage to run the application.
#
# @see http://label-schema.org for org.label-schema.* named labels.
#
#-------------------------------------------------------------------------------

FROM golang:1.13.0-alpine3.10 as builder

LABEL intermediate=true

ADD . /app
WORKDIR /app

# Install module dependencies, run test suite, and build executable.
RUN go mod vendor \
  && go test ./... -cover \
  && go build -v -mod=vendor -o ./go-link-check

FROM alpine:3.10.2

ARG BUILD_DATE=""
ARG VCS_REF=""
ARG VERSION=""

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="Go Link Check" \
      org.label-schema.description="Link check automation tool." \
      org.label-schema.url="https://github.com/dbtedman/go-link-check" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/dbtedman/go-link-check" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

WORKDIR /app

COPY --from=builder /app /app/

# Configure docker to execute our application on run.
ENTRYPOINT ./go-link-check
