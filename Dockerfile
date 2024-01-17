FROM golang:1.21-bookworm AS builder

ENV GOPRIVATE=*.cie.com
ENV GOINSECURE=*
RUN git config --global http.sslverify false
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.26.1/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf" && chmod +x /usr/local/bin/buf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
COPY . /build
WORKDIR /build
RUN make

FROM debian:bookworm
USER 0
COPY --from=builder /build/.bin/rssi-grpc /opt/rssi-grpc
USER 1000
WORKDIR /
CMD ["/opt/."]