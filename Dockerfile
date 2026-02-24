FROM alpine:latest as base
RUN apk update

FROM base as prepare
RUN apk add --no-cache g++ make libtool automake autoconf m4 git \
  && mkdir /opt/tsunami-udp-build \
  && cd /opt/tsunami-udp-build \
  && git clone https://git.code.sf.net/p/tsunami-udp/code_git .

FROM prepare as build
RUN mkdir /opt/tsunami-udp \
  && cd /opt/tsunami-udp-build/tsunami-udp \
  && make \
  && ./configure --prefix=/opt/tsunami-udp \
  && make install

FROM alpine:latest as runtime
COPY --from=build /opt/tsunami-udp /opt/tsunami-udp/
ADD . /opt/gsync
