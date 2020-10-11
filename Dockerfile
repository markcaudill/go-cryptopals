FROM golang:alpine as builder
MAINTAINER Mark Caudill <mark@mrkc.me>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	bash \
	ca-certificates

COPY . /go/src/github.com/markcaudill/go-cryptopals

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/markcaudill/go-cryptopals \
	&& make static \
	&& mv go-cryptopals /usr/bin/go-cryptopals \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM alpine:latest

COPY --from=builder /usr/bin/go-cryptopals /usr/bin/go-cryptopals

ENTRYPOINT [ "go-cryptopals" ]
CMD [ "--help" ]
