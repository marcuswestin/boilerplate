FROM golang

ADD . /go/src/examples/grpc-echo
RUN go install examples/grpc-echo/echo-server
ENTRYPOINT /go/bin/echo-server

EXPOSE 30001
