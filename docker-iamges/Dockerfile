FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

EXPOSE 9090

ADD ./prometheus-1.1.3.linux-amd64.tar.gz /usr/local

RUN mv /usr/local/prometheus-1.1.3.linux-amd64 /usr/local/prometheus && \
	ln -s /usr/local/prometheus/prometheus /usr/bin/prometheus && \
	ln -s /usr/local/prometheus/promtool /usr/bin/promtool && \
	cp /usr/local/prometheus/prometheus.yml /etc/prometheus.yml

ENTRYPOINT ["prometheus", "--config.file", "/etc/prometheus.yml"]