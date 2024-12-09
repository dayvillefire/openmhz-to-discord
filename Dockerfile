FROM busybox:glibc
WORKDIR /
ADD ./ca-certificates.crt /etc/ssl/certs/
ADD ./openmhz-to-discord /openmhz-to-discord
ENTRYPOINT ["/openmhz-to-discord"]
