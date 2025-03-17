# syntax=docker/dockerfile:1.4

FROM builder AS builder

FROM runner

COPY --from=builder /build/dist/cmd ./api

COPY config/local/api.yaml /config/local/api.yaml

ENTRYPOINT ["./api"]
