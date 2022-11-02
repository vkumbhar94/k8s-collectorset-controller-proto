FROM golang:1.19
WORKDIR $GOPATH/src/github.com/vkumbhar94/k8s-collectorset-controller-proto
RUN apt-get update
RUN apt-get -y install --no-install-recommends libarchive-tools
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip | bsdtar -xf - --strip-components=1 -C /bin bin/protoc \
  && chmod +x /bin/protoc
COPY ./proto ./proto
RUN mkdir api
RUN protoc -I proto proto/api.proto \
  --go_out=plugins=grpc:api --go_opt=paths=source_relative
RUN cp $GOPATH/src/github.com/vkumbhar94/k8s-collectorset-controller-proto/api/api.pb.go /tmp/
