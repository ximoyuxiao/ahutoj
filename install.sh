#! /bin/bash
sed -i 's/tencentyun/aliyun/g' /etc/apt/sources.list
sudo apt update
# 安装一系列工具
for pkg in "net-tools make flex g++ clang libmysqlclient-dev libmysql++-dev nginx mysql-server sysv-rc-conf pkg-config redis"
do
    echo "正在为您安装$pkg..."
	while ! apt-get install -y $pkg 
	do
		echo "Network fail, retry... you might want to change another apt source for install"
	done
done

# 安装go环境
which go
if [ $? -eq 1 ]
    then
        echo "正在为您安装go..."
    else
        echo "go已经安装"
        sudo add-apt-repository ppa:longsleep/golang-backports
        sudo apt update
        sudo apt install golang-go
        go env -w GO111MODULE="on"
        go env -w GOPROXY="https://goproxy.cn,direct"
        go env -w GOSUMDB=off
fi
# 安装air
which air
if [ $? -eq 0 ]
then
    echo "air 已经安装"
else
    echo "给您安装go air..."
    go install github.com/cosmtrek/air@latest
fi

# 建立数据库
USER=`sudo cat /etc/mysql/debian.cnf |grep user|head -1|awk  '{print $3}'`
PASSWORD=`sudo cat /etc/mysql/debian.cnf |grep password|head -1|awk  '{print $3}'`
CPU=`grep "cpu cores" /proc/cpuinfo |head -1|awk '{print $4}'`
mysql -h localhost -u$USER -p$PASSWORD < ./doc/oj.sql
echo "insert into jol.privilege values('admin','administrator','true','N');"|mysql -h localhost -u$USER -p$PASSWORD 
cp ./config.yaml.bak ./config.yaml
make build
echo "username:$USER"
echo "password:$PASSWORD"