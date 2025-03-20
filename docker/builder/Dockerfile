# syntax=docker/dockerfile:1.4
FROM schema as schema

FROM golang:1.23.0 AS builder

ARG GO_RACE_DETECTOR

ENV GO_RACE_DETECTOR=${GO_RACE_DETECTOR} GORACE="abort_on_error=1 halt_on_error=1"

WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum Makefile ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && make tools

COPY sqlc.yaml ./
COPY ./openapi ./openapi
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./db ./db

# Build the Go app
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/build \
    make go

FROM golang:1.23.0 AS runner

WORKDIR /
