pkger -o web -include /web/static/ -include /web/templates/ && \
CC=/usr/local/bin/musl-gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fiber-pongo2-pkger -a -ldflags '-extldflags "-static" -s -w'  . && \
upx fiber-pongo2-pkger
