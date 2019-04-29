FROM debian:stretch
MAINTAINER Sean Condon <sean@opennetworking.org>, Adib Rastegarnia <adib.rastegarnia@gmail.com> 
LABEL description="Builds a gNOI simulator on a Debian distribution"

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -qy --no-install-recommends \
      apt-utils \
      ca-certificates \
      software-properties-common \
    && echo "deb http://deb.debian.org/debian stretch-backports main" >  /etc/apt/sources.list.d/backports.list \
    && apt-get update

RUN DEBIAN_FRONTEND=noninteractive apt-get install -qy --no-install-recommends \
    fping \
    git \
    golang-1.11-go \
    iproute2 \
    iputils-ping \
    net-tools \
    netcat-openbsd \
    openssh-client \
    psmisc \
    sudo \
    vim

ENV HOME=/home/gniosim
RUN mkdir $HOME
WORKDIR $HOME

ENV GOROOT=/usr/lib/go-1.11
ENV GOBIN=$GOROOT/bin
ENV PATH=$GOBIN:${PATH}
ENV GOPATH=$HOME/go

RUN mkdir -p $GOPATH \
    && go get -u \
            github.com/google/gnxi/gnoi_target \ 
            github.com/google/gnxi/gnoi_cert 
RUN go install -v \
            github.com/google/gnxi/gnoi_target \
            github.com/google/gnxi/gnoi_cert 
         

ENV GNOI_TARGET=localhost
ENV GNOI_PORT=50002

COPY ./certs certs

CMD \
  IPADDR=`ip route get 1.2.3.4 | grep dev | awk '{print $7}'` && \
  if [ "$GNOI_TARGET" != "localhost" ]; then \
      certs/generate_certs.sh $GNOI_TARGET > /dev/null 2>&1; \
      echo "Please add '"$IPADDR" "$GNOI_TARGET"' to /etc/hosts and access with gNOI client at "$GNOI_TARGET":"$GNOI_PORT; \
  else \
      echo "gNOI running on localhost:"$GNOI_PORT; \
  fi && \
  gnoi_target \
  -bind_address :$GNOI_PORT \
  -alsologtostderr \



