# docker 17.05+ is required for multistage build

FROM golang:1.8 as build
COPY main.go /main.go

RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -s" -o showcase-app main.go

FROM scratch
LABEL maintainer "Jan Garaj <jan.garaj@gmail.com>"

ENV \
  PORT=:443 \
  READ_TIMEOUT=60 \
  WRITE_TIMEOUT=600 \
  TLS_CERT=/certs/cert.pem \
  TLS_KEY=/certs/key.pem \
  TLS_MIN_VERSION=VersionTLS12 \
  TLS_CIPHER=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256

WORKDIR /
ENTRYPOINT ["/showcase-app"]
COPY files /
COPY --from=build /showcase-app /

