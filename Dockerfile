FROM golang:1.16.0-alpine AS go-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -o /go/bin/lazarette

FROM mhart/alpine-node:15.11.0 AS svelte-builder

WORKDIR /app
COPY frontend ./
RUN npm ci && npm run build

FROM alpine:3.13.2

LABEL org.label-schema.description="Kubernetes Volumes Explorer" \
    org.label-schema.name="lazarette" \
    org.label-schema.url="https://github.com/acim/lazarette/blob/master/README.md" \
    org.label-schema.vendor="ablab.io"

WORKDIR /app
COPY --from=go-builder /go/bin/lazarette /usr/bin/lazarette
COPY --from=svelte-builder /app/public public/

EXPOSE 3000

USER 65534

ENTRYPOINT ["/usr/bin/lazarette"]
