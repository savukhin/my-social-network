# Golang
FROM golang:1.18.3-alpine
# RUN apk add --no-cache su-exec
RUN set -ex && apk --no-cache add sudo

WORKDIR /usr/src/app/server
COPY . .

RUN export GOPATH=/usr/src/app/server
RUN export PATH=$GOPATH/bin:$PATH

RUN sh ./build.sh
EXPOSE 4201 

ENTRYPOINT ["sh", "./run.sh"]
