#! /bin/bash

cd /root/Desktop/code/src/k8sBeegoDemo
time=$(date "+%m.%d")
docker build -t beego:v$time . && docker tag beego:v$time xxz199439/beegodemo:latest && docker push xxz199439/beegodemo:latest && docker rmi -f $(docker images | grep beego |grep -v grep|awk '{print $3}')








