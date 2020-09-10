FROM golang:1.15.2-alpine AS go-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o /go/bin/kve

FROM mhart/alpine-node:14 AS svelte-builder

WORKDIR /app
COPY frontend ./
RUN npm ci && npm run build

FROM alpine:3.12.0

LABEL org.label-schema.description="Kubernetes Volumes Explorer" \
    org.label-schema.name="kve" \
    org.label-schema.url="https://github.com/acim/kve/blob/master/README.md" \
    org.label-schema.vendor="ablab.io"

WORKDIR /app
COPY --from=go-builder /go/bin/kve /usr/bin/kve
COPY --from=svelte-builder /app/public public/

EXPOSE 3000

USER 65534

ENTRYPOINT ["/usr/bin/kve"]
