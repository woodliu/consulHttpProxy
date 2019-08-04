FROM docker.io/centos:latest
COPY consulProxyServer /usr/local/bin/
ENTRYPOINT ["consulProxyServer"]
