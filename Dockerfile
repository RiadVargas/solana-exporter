FROM golang as builder

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/solana-exporter /

ENTRYPOINT ["/solana-exporter"]