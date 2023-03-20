#! /bin/bash
sed -i 's/tencentyun/aliyun/g' /etc/apt/sources.list
sudo apt update
# 安装一系列工具
for pkg in net-tools make flex g++ clang libmysqlclient-dev libmysql++-dev nginx mysql-server pkg-config redis libhiredis-dev cmake
do
    echo "正在为您安装$pkg..."
	if ! apt-get install -y $pkg
    then
		echo "Network fail, retry... you might want to change another apt source for install"
        exit 1
	fi
done

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

# 安装所有软件
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