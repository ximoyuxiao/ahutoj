#!/bin/sh
update=/root/update
WEBSERVER=$update/ahutoj/
WEBCLIENT=$update/AhutOjVue/
mkdir $update
cd $update
git clone git@github.com:199094212/ahutoj.git
cd $WEBSERVER
make build
cd core 
make clean
make
judgedID=ps -ef|head -n -1|grep judged|awk '{print $2}'
kill -9 $judgedID
`./judged`
cd ..
nohup ./tmp/bin/main ./config.yaml &
cd ..
git clone git@github.com:cz2542079957/AhutOjVue.git
cd $WEBCLIENT
git checkout -b release && git pull origin release
tar -zxvf ./dist.rar /root/AhutOjVue/dist
echo 'update success'