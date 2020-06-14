FROM golang:1.14-buster as build-env

ENV GOPROXY=https://proxy.golang.org

RUN apt-get -y update \
    && apt-get install -y make git upx wget curl tar musl* \
    && go get -u github.com/swaggo/swag/cmd/swag

ARG BUILD_DATE
ARG VCS_REF

WORKDIR /opt/src

ADD . /opt/src

RUN export PATH=/go/bin:$PATH && \
    swag init --generalInfo pkg/routers/routers.go pkg/routers && \
    CC=/usr/local/bin/musl-gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/fiber-pongo2-pkger  -a -ldflags '-extldflags "-static" -s -w'  .

RUN upx /opt/src/build/fiber-pongo2-pkger

FROM scratch
# FROM gcr.io/distroless/base
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /opt/src/build/fiber-pongo2-pkger /
COPY --from=build-env /opt/src/docs /

ENTRYPOINT ["/fiber-pongo2-pkger"]
