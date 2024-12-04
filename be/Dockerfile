ARG GO_IMAGE=golang:1.22.6-bullseye
ARG DEBIAN_IMAGE=debian:11-slim

##
## Build
##
FROM ${GO_IMAGE} AS build-base

# Add key for private repository
ARG GITHUB_PASSWORD
ENV GITHUB_PASSWORD=$GITHUB_PASSWORD

RUN git config \
    --global \
    url."https://${GITHUB_PASSWORD}@github.com".insteadOf \
    "https://github.com"

WORKDIR /app

RUN go env -w GOMODCACHE=/root/.cache/go-build
COPY go.mod go.sum ./
COPY pkg ./pkg
RUN --mount=type=cache,target=/root/.cache/go-build go mod download -x

COPY cmd ./cmd
COPY internal ./internal

FROM build-base AS build-api
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /go/bin/app ./cmd/promotion-api

FROM build-base AS build-all
COPY --from=build-promotion-api /go/bin/app /go/bin/app-promotion-api

##
## Dev
##
FROM build-base AS dev

# install air for hot reload
RUN curl -sSfL --proto "=https" https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./.air.toml .
COPY . .

CMD ["air", "-build.cmd", "go build -o ./tmp/main ./cmd"]

##
## Deploy
##
FROM ${DEBIAN_IMAGE} AS deploy-base

WORKDIR /app

RUN groupadd -r nonroot && useradd -r -g nonroot nonroot
# NOTE: need wget for getting Amazon Root CA for TLS signing
RUN apt update && apt install -y jq wget && rm -rf /var/lib/apt/lists/*
RUN wget https://www.amazontrust.com/repository/AmazonRootCA1.pem -O /etc/ssl/certs/ca-certificates.crt && update-ca-certificates

USER nonroot:nonroot

COPY --chown=nonroot:nonroot asset ./asset

EXPOSE 8080

FROM deploy-base AS deploy-promotion-api
COPY --chown=nonroot:nonroot conf/promotion-api ./conf/promotion-api
COPY --chown=nonroot:nonroot internal/migration/ ./internal/migration/
COPY --from=build-promotion-api --chown=nonroot:nonroot /go/bin/app ./app
CMD ["./app"]