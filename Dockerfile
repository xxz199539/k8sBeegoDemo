FROM golang:alpine AS builder
WORKDIR /root/Desktop/code/src/k8sBeegoDemo
COPY . .
ENV GOPATH /root/Desktop/code
#RUN sh govendor.sh
#RUN sh env.sh
RUN  pwd
RUN go build main.go
ENTRYPOINT ["./main"]