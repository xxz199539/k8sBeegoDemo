FROM golang:alpine AS builder
WORKDIR /root/Desktop/code/src/k8sBeegoDemo

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY . .
ENV GOPATH /root/Desktop/code
#RUN sh govendor.sh
#RUN sh env.sh
RUN pwd
RUN go build /root/Desktop/code/src/k8sBeegoDemo/main.go
CMD ["./main"]
