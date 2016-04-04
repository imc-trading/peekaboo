# Dev. environment
# docker run -it --rm -v $PWD:/peekaboo -v /var/run/docker.sock:/var/run/docker.sock -p 5050:5050 <image id>

FROM centos:centos7

RUN set -ex ;\
    yum install -y go vim-enhanced net-tools lsb lvm2 docker git ;\
    yum clean all

ENV GOPATH=/root/go
ENV PATH=${PATH}:${GOPATH}/bin
ENV PROJECT=/peekaboo

RUN set -ex ;\
    go get github.com/constabulary/gb/... ;\
    mv ${GOPATH}/bin/gb /usr/local/bin/gb ;\
    mv ${GOPATH}/bin/gb-vendor /usr/local/bin/gb-vendor ;\
    mkdir ${PROJECT}

EXPOSE 5050

WORKDIR ${PROJECT}
ENTRYPOINT /bin/bash
