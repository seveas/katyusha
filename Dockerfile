FROM golang:1.16

WORKDIR /katyusha
COPY . .
RUN touch provider/plugin/common/plugin.pb.go provider/plugin/common/plugin_grpc.pb.go
RUN make
RUN make -C integration/pki install-ca
