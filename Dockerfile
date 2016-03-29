FROM alpine:3.3
MAINTAINER chengtiesheng <chengtiesheng@huawei.com>

RUN apk add --no-cache \
		bash \
		iptables \
		ca-certificates \
		e2fsprogs              

ADD https://get.docker.com/builds/Linux/x86_64/docker-1.9.1 /usr/bin/docker
ADD *.sh /
COPY builder /usr/local/bin/builder
RUN chmod +x /usr/bin/docker /usr/local/bin/* && rm -fr /var/lib/docker/*

CMD ["/run.sh"]
