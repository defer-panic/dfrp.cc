FROM golang:1.19 as BUILD

WORKDIR /build

COPY . .

RUN make build

FROM alpine:3.16

COPY --from=BUILD /build/bin/ /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/url-shortener-api"]
