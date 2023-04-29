#! /bin/bash
sed -i 's/tencentyun/aliyun/g' /etc/apt/sources.list
sudo apt update
# 安装一系列工具
for pkg in net-tools make flex g++ clang libmysqlclient-dev libmysql++-dev nginx mysql-server pkg-config redis libhiredis-dev cmake erlang-nox
do
    echo "正在为您安装$pkg..."
	if ! apt-get install -y $pkg
    then
		echo "Network fail, retry... you might want to change another apt source for install"
        exit 1
	fi
done
which rabbitmq-server
if [$? -eq 1]
    then
        echo "正在为您安装rabbitMQ"
        wget -O- https://www.rabbitmq.com/rabbitmq-release-signing-key.asc | sudo apt-key add -
        sudo apt-get update
        sudo apt-get install -y rabbitmq-server  #安装成功自动启动    
        rabbitmq-plugins enable rabbitmq_management   # 启用插件
        service rabbitmq-server restart    # 重启
        rabbitmqctl add_user ahutoj 2019ahut   # 增加普通用户
        rabbitmqctl set_user_tags ahutoj administrator    # 给普通用户分配管理员角色
    else
        echo "rabbitMQ已经安装"
# 安装go环境
which go
if [ $? -eq 1 ]
    then
        echo "正在为您安装go..."
        sudo add-apt-repository ppa:longsleep/golang-backports
        sudo apt update
        sudo apt install -y golang-go
        go env -w GO111MODULE="on"
        go env -w GOPROXY="https://goproxy.cn,direct"
        go env -w GOSUMDB=off
        # 这一块  没办法自动获取  后面要改
        export GOPATH=`go env GOPATH`
        export PATH=$PATH:/usr/bin/go:${GOPATH}:${GOPATH}/bin
    else
        echo "go已经安装"
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

echo "正在安装nlohmannjson-dev"
wget https://github.com/nlohmann/json/archive/refs/tags/v3.11.2.tar.gz
tar -zxvf v3.11.2.tar.gz
cd json-3.11.2
mkdir build && cd build
cmake ..
make
make install
cd ..
cd ..

echo "正在安装libamqp-dev"
wget https://github.com/alanxz/rabbitmq-c/archive/refs/tags/v0.9.0.tar.gz
tar -zxvf v0.9.0.tar.gz
cd rabbitmq-c-0.9.0
mkdir build && cd build
cmkae ..
make 
make install

cd ..
cd ..
# 安装所有软件
ldconfig
make
make install
# 建立数据库
USER=`sudo cat /etc/mysql/debian.cnf |grep user|head -1|awk  '{print $3}'`
PASSWORD=`sudo cat /etc/mysql/debian.cnf |grep password|head -1|awk  '{print $3}'`
CPU=`grep "cpu cores" /proc/cpuinfo |head -1|awk '{print $4}'`
mysql -h localhost -u$USER -p$PASSWORD < ./doc/oj.sql
# echo "insert into ahutoj.Perrmission values('admin','administrator','true','N');"|mysql -h localhost -u$USER -p$PASSWORD
sed -i 's/skip-external-locking/# skip-external-locking/g' /etc/mysql/mysql.conf.d/mysqld.cnf
sed -i 's/bind-address            = 127.0.0.1/bind-address            = 0.0.0.0/g' /etc/mysql/mysql.conf.d/mysqld.cnf
systemctl restart mysql
cp ./config.yaml.bak ./config.yaml
redis-cli < ./doc/redis.in
make build
cd core
make all
cd ..
echo "username:$USER"
echo "password:$PASSWORD"

git clone https://github.com/nlohmann/json.git
cd json
mkdir build
cd build/
cmake ..
make
make install
installNPM(){
    pwd=`pwd`
    ln -s $pwd/node-v16.3.0-linux-x64/bin/node /usr/local/bin/node
    ln -s $pwd/node-v16.3.0-linux-x64/bin/npm /usr/local/bin/npm
    npm i vue-tsc -D
    npm install -g vite
}