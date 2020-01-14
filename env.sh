#! /bin/bash

export JAVA_HOME="/root/Downloads/jdk8"
export JRE_HOME="/root/Downloads/jdk8/jre"
export CLASSPATH="$JAVA_HOME/lib:$JRE_HOME/lib"
export GOROOT="/opt/go"
export GOTOOLS="$GOROOT/pkg/tool"
export GOARCH=386
export GOOS="linux"
export GOPATH="/root/Desktop/code"
export M2_HOME="/root/Desktop/maven3"
export CLASSPATH="$CLASSPATH:$M2_HOME/Lib"
export PATH="$PATH:$JAVA_HOME/bin:$GOROOT/bin:$M2_HOME/bin"
mv /etc/apt/sources.list /etc/apt/sources.list.bak

Codename=$( (lsb_release -a)|awk '{print $2}'|tail -n 1 )


 
echo "\

deb http://mirrors.aliyun.com/ubuntu/ $Codename main multiverse restricted universe

deb http://mirrors.aliyun.com/ubuntu/ $Codename-backports main multiverse restricted universe

deb http://mirrors.aliyun.com/ubuntu/ $Codename-proposed main multiverse restricted universe

deb http://mirrors.aliyun.com/ubuntu/ $Codename-security main multiverse restricted universe

deb http://mirrors.aliyun.com/ubuntu/ $Codename-updates main multiverse restricted universe

deb-src http://mirrors.aliyun.com/ubuntu/ $Codename main multiverse restricted universe

deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-backports main multiverse restricted universe

deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-proposed main multiverse restricted universe

deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-security main multiverse restricted universe

deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-updates main multiverse restricted universe ">/etc/apt/sources.list
apt-get update
