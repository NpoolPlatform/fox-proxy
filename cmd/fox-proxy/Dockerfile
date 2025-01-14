FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

COPY FoxProxy.viper.yaml /etc/FoxProxy/FoxProxy.viper.yaml
COPY fox-proxy /usr/local/bin/fox-proxy

CMD ["fox-proxy", "run"]
