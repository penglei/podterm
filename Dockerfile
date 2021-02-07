FROM alpine

ENTRYPOINT ["/usr/local/bin/entrypoint"]

RUN apk add jq gettext curl openssl bash

ADD entrypoint /usr/local/bin/
