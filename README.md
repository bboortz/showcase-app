# Showcase App

Showcase App, which prints request headers.

## Docker

5MB Docker image [jangaraj/showcase-app](https://hub.docker.com/r/jangaraj/showcase-app/)
is available. Example of Docker run with custom certificates:

```bash
docker run \
  --name showcase-app \
  -p 443:443 \
  -v $PWD/mycerts:/mycerts \
  -e PORT=:443 \
  -e TLS_CERT=/mycerts/mycert.pem \
  -e TLS_KEY=/mycerts/mykey.pem \
  -e TLS_CIPHER=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 \
  -e TLS_MIN_VERSION=VersionTLS12 \
  jangaraj/showcase-app
```

## Env variables

| Env variable | Default value | Description |
| :----------: | :-----------: | :---------: |
| PORT | *443* | TCP port of app |
| READ_TIMEOUT | 60 | Time from when the connection is accepted to when the request body is fully read (if you do read the body, otherwise to the end of the headers). |
| WRITE_TIMEOUT | 600 |  Time from the end of the request header read to the end of the response write. |
| TLS_CERT | */certs/cert.pem* | TLS certificate |
| TLS_KEY | */certs/key.pem* | TLS key |
| TLS_MIN_VERSION | *VersionTLS12* | Minimum SSL/TLS version that is acceptable.<br>[Possible values](https://golang.org/pkg/crypto/tls/):<br>*VersionTLS10, VersionTLS11, VersionTLS12, VersionSSL30* |
| TLS_CIPHER | *TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256* | Supported cipher suite(s) - use *\|* as a separator if you need to specify more suites.<br>[Possible values](https://golang.org/pkg/crypto/tls/):<br>*TLS_RSA_WITH_RC4_128_SHA, TLS_RSA_WITH_3DES_EDE_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_RC4_128_SHA, TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, TLS_FALLBACK_SCSV* |

## Debugging

Please check container logs. They may contain usefull information. Example:

```bash
$ docker log showcase-app
2017/08/01 12:05:52 Starting showcase-app on the port :443
2017/08/01 12:05:52 Env variable: PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
2017/08/01 12:05:52 Env variable: HOSTNAME=215fd16c8544
2017/08/01 12:05:52 Env variable: PORT=:443
2017/08/01 12:05:52 Env variable: TLS_CERT=/mycerts/mycert.pem
2017/08/01 12:05:52 Env variable: TLS_KEY=/mycerts/mykey.pem
2017/08/01 12:05:52 Env variable: TLS_CIPHER=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
2017/08/01 12:05:52 Env variable: TLS_MIN_VERSION=VersionTLS12
2017/08/01 12:05:52 Env variable: HOME=/
2017/08/01 12:06:02 Processing request: GET /public/ HTTP/2.0
2017/08/01 12:06:12 Processing request: POST /private/ HTTP/2.0
```

