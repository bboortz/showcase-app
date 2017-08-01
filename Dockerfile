# docker 17.05+ is required for multistage build

FROM golang:1.8 as build
COPY main.go /

RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -s" -o showcase-app main.go

FROM scratch
LABEL maintainer "Jan Garaj <jan.garaj@gmail.com>"

# A list of cipher suite IDs:
# https://golang.org/pkg/crypto/tls/

ENV \
  PORT=:443 \
  TLS_CERT=/certs/cert.pem \
  TLS_KEY=/certs/key.pem \
  TLS_MIN_VERSION=VersionTLS12 \
  TLS_CIPHER=TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256

WORKDIR /
ENTRYPOINT ["/showcase-app"]
COPY certs /certs
COPY --from=build /showcase-app /

