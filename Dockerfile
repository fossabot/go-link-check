#-------------------------------------------------------------------------------
#
# Docker Configuration, https://docs.docker.com/engine/reference/builder/
#
# Includes "builder" stage to compile and test application and slimmed down
# "runner" stage to run the application.
#
#-------------------------------------------------------------------------------

FROM golang:1.13.0-alpine3.10 as builder

LABEL go-link-check=true

ADD . /app
WORKDIR /app

# Install module dependencies, run test suite, and build executable.
RUN go mod vendor \
  && go test ./... -cover \
  && go build -v -mod=vendor -o ./go-link-check

#-------------------------------------------------------------------------------

FROM alpine:3.10.2 as runner

LABEL go-link-check=true

WORKDIR /app

COPY --from=builder /app /app/

# Configure docker to execute our application on run.
ENTRYPOINT ./go-link-check
