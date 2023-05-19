# first stage
FROM golang:1.17 as build
WORKDIR /src

RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/win2key/jvpn.git .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /output/jvpn
RUN cp openconnect-expect.sh /output/openconnect-expect.sh && chmod +x /output/openconnect-expect.sh

# second stage
FROM alpine:3.17 as build2
WORKDIR /src

RUN apk update && apk add --no-cache \
    build-base \
    libxml2-dev \
    gnutls-dev \
    libnl3-dev \
    libev-dev \
    iptables-dev \
    uwsgi-tuntap \
    gettext \
    expect \
    git \
    autoconf \
    automake \
    libtool \
    lz4-dev

RUN mkdir -p /etc/vpnc && wget -c -O /etc/vpnc/vpnc-script https://gitlab.com/openconnect/vpnc-scripts/-/raw/master/vpnc-script && \
    chmod +x /etc/vpnc/vpnc-script

RUN git clone https://gitlab.com/openconnect/openconnect.git && \
    cd openconnect && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install 

RUN tar czhf opn.tar.gz $(ldd /usr/local/sbin/openconnect | sed -E 's/^[[:space:]]*//;s/.*=> | \(.*\)//g' | sort -u | tr '\n' ' ') /usr/local/sbin/openconnect \
    /etc/vpnc/vpnc-script $(ldd /usr/bin/expect | sed -E 's/^[[:space:]]*//;s/.*=> | \(.*\)//g' | sort -u | tr '\n' ' ') /usr/bin/expect /usr/lib/tcl8.6/init.tcl 

# final stage
FROM alpine:3.17
WORKDIR /app

COPY --from=build /output/jvpn /app/jvpn
COPY --from=build /output/openconnect-expect.sh /app/openconnect-expect.sh

COPY --from=build2 /src/opn.tar.gz /app/opn.tar.gz
RUN tar xzf opn.tar.gz -C / && rm -f /app/opn.tar.gz

RUN echo tun >> /etc/modules

EXPOSE 8080
EXPOSE 8081

CMD ["/app/jvpn"]