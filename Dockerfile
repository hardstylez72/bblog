FROM alpine:3.12.0

CMD mkdir /opt/bblog
WORKDIR /opt/bblog
COPY /cmd/server/main .

CMD ["/opt/bblog/main", "-config", "/etc/app/config.yaml"]