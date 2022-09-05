# Golang
FROM golang:1.18.3-alpine

WORKDIR $GOPATH
COPY . .

RUN sh ./build.sh
EXPOSE 4201 

ENTRYPOINT ["sh", "./run.sh"]
