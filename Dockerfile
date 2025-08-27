FROM ubuntu:xenial

ENV INSTALL_DEPS \
  ca-certificates \
  gcc \
  git \
  make \
  software-properties-common \
  unzip \
  wget \
  ssh

RUN apt-get update \
  && apt-get install -y -q --no-install-recommends ${INSTALL_DEPS} \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# protoc
ENV PROTOC_VER=3.15.0
ENV PROTOC_REL=protoc-"${PROTOC_VER}"-linux-x86_64.zip
RUN wget https://github.com/google/protobuf/releases/download/v"${PROTOC_VER}/${PROTOC_REL}" \
  && unzip ${PROTOC_REL} -d protoc \
  && mv protoc /usr/local \
  && ln -s /usr/local/protoc/bin/protoc /usr/local/bin

# go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GORELEASE go1.16.linux-amd64.tar.gz
RUN wget -q https://dl.google.com/go/$GORELEASE \
  && tar -C $(dirname $GOROOT) -xzf $GORELEASE \
  && rm $GORELEASE \
  && mkdir -p $GOPATH/{src,bin,pkg}

# protoc-gen-go
ENV PGG_VER=v1.26.0
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PGG_VER}

# protoc-gen-grpc-go
ENV PGG_VER=v1.0.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PGG_VER}

# protoc-gen-graphql
RUN go install github.com/kitt-technology/protoc-gen-graphql@latest

# protoc-gen-graphql
RUN go install github.com/kitt-technology/protos-common/common@latest

WORKDIR /go/src/github.com/kitt-technology/protoc-gen-graphql
