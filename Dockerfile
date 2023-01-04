FROM alpine:latest
MAINTAINER sandy <sandy@ever.finance>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /permastats

COPY cmd/permastats /permastats/permastats
EXPOSE 8080

ENTRYPOINT [ "/permastats/permastats" ]