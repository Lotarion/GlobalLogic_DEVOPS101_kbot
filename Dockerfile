FROM golang AS builder

WORKDIR /go/src/app
COPY . .
RUN make get_deps
RUN make build

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/bin/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot"]
