FROM golang:1.15.0-buster

LABEL org.label-schema.description="kve development image" \
    org.label-schema.name="ablab/kve-dev" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.vendor="ablab.io"

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - && \
    apt-get install -y nodejs

RUN go get -u github.com/cespare/reflex

COPY reflex.conf /usr/local/etc/
COPY build.sh start.sh /usr/local/bin/

WORKDIR /app

VOLUME /go
VOLUME /app/frontend/node_modules

CMD ["/usr/local/bin/start.sh"]

EXPOSE 3000 35729
