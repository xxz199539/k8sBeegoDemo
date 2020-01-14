#! /bin/bash

echo "Stopping old contain"
for contain in `docker ps -a | grep ./main |grep -v grep|awk '{print $1}'`
do 
    docker stop $contain && docker rm $contain
done
for mysqlContain in `docker ps -a | grep mysql:5.7 |grep -v grep|awk '{print $1}'`
do 
    docker start $mysqlContain
done
for noneContain in `docker images | grep none |grep -v grep|awk '{print $3}'`
do 
    if [ ! $(docker ps -a | grep $noneContain |grep -v grep|awk '{print $1}') ];then
       echo "no contain need remove"
    else
        docker rm $(docker ps -a | grep $noneContain |grep -v grep|awk '{print $1}')
    fi
    docker rmi $noneContain
done
set timeout 1
spawn sudo docker login --username=550276107@qq.com registry.cn-chengdu.aliyuncs.com
expect "*password:"
send "xxz199439/r"
interact
sh env.sh
bash dockerBuild.sh








