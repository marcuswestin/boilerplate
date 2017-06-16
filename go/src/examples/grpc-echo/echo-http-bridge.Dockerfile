FROM golang

ADD . /go/src/examples/grpc-echo
RUN go install examples/grpc-echo/echo-http-bridge
ENTRYPOINT /go/bin/echo-http-bridge

EXPOSE 30002
