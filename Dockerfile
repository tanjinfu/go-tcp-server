FROM golang:1.6
EXPOSE 31069

MAINTAINER Tan Jinfu

ADD tcp-server.go /var/server/server.go

CMD ["go", "run", "/var/server/server.go"]