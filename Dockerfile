# docker run -it --rm -v $GOPATH:/go -v /var/run/docker.sock:/var/run/docker.sock -p 5050:5050 <id>

FROM centos:centos7

RUN set -ex ;\
    yum install -y go vim-enhanced net-tools lsb docker git wget ethtool ;\
    yum clean all

ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin
ENV PROJECT=github.com/imc-trading/peekaboo

# Add mock binaries
COPY mock/ipmitool /usr/local/bin/ipmitool
COPY mock/pvs /usr/local/bin/pvs
COPY mock/lvs /usr/local/bin/lvs
COPY mock/vgs /usr/local/bin/vgs
COPY mock/onload /usr/local/bin/onload
COPY mock/sfkey /usr/local/bin/sfkey
COPY mock/sfctool /usr/local/bin/sfctool

RUN chmod +x /usr/local/bin/ipmitool \
    /usr/local/bin/pvs \
    /usr/local/bin/lvs \
    /usr/local/bin/vgs \
    /usr/local/bin/onload \
    /usr/local/bin/sfkey \
    /usr/local/bin/sfctool

# Add mock files
COPY mock/config /config
RUN mkdir /boot ;\
    mv /config /boot/config-$(uname -r)

EXPOSE 5050

WORKDIR ${GOPATH}/src/${PROJECT}
ENTRYPOINT /bin/bash
