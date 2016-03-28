# Self containerd Dev. environment
# docker run -it --rm -v $HOME/go:/root/go -v /var/run/docker.sock:/var/run/docker.sock -p 5050:5050 <image id>

FROM centos:centos7

RUN set -ex ;\
    yum install -y go vim-enhanced net-tools lsb lvm2 docker ;\
    yum clean all

ENV GOPATH=/root/go
ENV PATH=${PATH}:${GOPATH}/bin
ENV PROJECT=${GOPATH}/src/github.com/imc-trading/peekaboo

RUN go get github.com/constabulary/gb/...

EXPOSE 5050

WORKDIR ${PROJECT}
ENTRYPOINT /bin/bash
