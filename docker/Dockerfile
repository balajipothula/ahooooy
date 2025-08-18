FROM alpine:3.22.1 AS nginx1.28.0

LABEL \
  # Authors or maintainers of the image.
  org.opencontainers.image.authors="Balaji Pothula <balan.pothula@gmail.com>" \
  # Image reference of the image.
  org.opencontainers.image.base.name="alpine:3.22.1" \
  # Title or name for the image.
  org.opencontainers.image.title="nginx:1.28.0" \
  # Version of the packaged software.
  org.opencontainers.image.version="1.28.0" \
  # Description or summary of the image.
  org.opencontainers.image.description="NGINX 1.28.0 - Reverse Proxy Server" \
  # Name of the distributing entity, organization or individual.
  org.opencontainers.image.vendor="Blitz Systems Pvt. Ltd. <https://blitzsystems.in>"

RUN \
  # Exit immediately if any command fails.
  set -e; \
  # Install the required dependencies and packages.
  apk add --no-progress --no-cache \
    nginx-mod-http-lua \
    nginx-mod-http-geoip \
    nginx-mod-stream-geoip \
    lua5.1-maxminddb \
    lua5.1-socket \
    sqlite-libs \
    redis && \
  # Remove Debian based packages.
  rm -rf /tmp/* /var/tmp/* /var/cache/apk/* /var/cache/distfiles/*

# Mount point for NGINX configuration files.
VOLUME ["/nginx/conf"]

# Expose port 443 for incoming HTTP traffic.
# Container will listen on port 443 using TCP.
EXPOSE 443/tcp

# Set the entrypoint to start Nginx with the specified configuration file.
ENTRYPOINT ["nginx", "-c", "/nginx/conf/nginx.conf", "-g", "daemon off;"]
