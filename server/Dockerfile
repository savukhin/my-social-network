# Golang
FROM golang:1.18.3-alpine
# RUN apk add --no-cache su-exec
RUN set -ex && apk --no-cache add sudo

WORKDIR $GOPATH
COPY . .

RUN sh ./build.sh
EXPOSE 4201 

ENTRYPOINT ["sh", "./run.sh"]
