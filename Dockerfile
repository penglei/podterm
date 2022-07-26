FROM alpine

ENTRYPOINT ["/usr/local/bin/entrypoint"]

RUN apk add jq gettext curl openssl bash openssh iproute2

ADD entrypoint /usr/local/bin/
