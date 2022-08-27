# Golang
FROM golang:1.18.3-alpine
# RUN apk add --no-cache su-exec
RUN set -ex && apk --no-cache add sudo

WORKDIR /usr/src/app/server
COPY . .
# RUN chmod +x build.sh
# RUN chmod +x run.sh

RUN export GOPATH=/usr/src/app/server
RUN export PATH=$GOPATH/bin:$PATH

# WORKDIR /usr/src/app/server/src/api
# RUN go mod download
# WORKDIR /usr/src/app/server/src/api 
# RUN go install api 
# WORKDIR /usr/src/app/server
RUN sh ./build.sh
EXPOSE 4201 

ENTRYPOINT ["sh", "./run.sh"]

# USER root:root
# CMD ["./bin/api"]
# CMD ["bash", "./run.bash"]
# CMD [ "ls",  "/usr/src/app/server/bin" ]
# CMD [ "ls",  "/usr/src/app/server/" ]
# ENTRYPOINT [ "ls",  "/usr/src/app/server/pkg" ]


# PosgreSQL
# FROM postgres:14.0-alpine

# USER postgres

# RUN chmod 0700 /var/lib/postgresql/data &&\
#     initdb /var/lib/postgresql/data &&\
#     echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf &&\
#     echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf &&\
#     pg_ctl start &&\
#     psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'socialNetwork'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE socialNetwork" &&\
#     psql -c "ALTER USER postgres WITH ENCRYPTED PASSWORD 'admin';"

# EXPOSE 5432
