# Atlas Dockerfile
FROM alpine:latest
RUN apk add go
COPY ./api /usr/local/bin/api
EXPOSE 8800/tcp
EXPOSE 80/tcp
CMD ["/usr/local/bin/api", "-f /data/config.json"]
