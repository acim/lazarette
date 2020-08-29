FROM golang:1.15.0-buster

LABEL org.label-schema.description="Docker image for autorecompiling and autorestarting Golang server" \
    org.label-schema.name="acim/go-reflex" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.url="https://github.com/acim/go-reflex/blob/master/README.md" \
    org.label-schema.vendor="ablab.io"

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - && \
    apt-get install -y nodejs

RUN go get -u github.com/cespare/reflex

COPY reflex.conf /usr/local/etc/
COPY build.sh /usr/local/bin/

WORKDIR /app

VOLUME /go

CMD ["reflex", "-d", "none", "-c", "/usr/local/etc/reflex.conf"]

EXPOSE 3000
