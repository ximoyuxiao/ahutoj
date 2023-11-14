#!/bin/sh
update=/home/ubuntu/
WEBSERVER=$update/ahutoj/
mkdir $update
cd $update
git config pull.rebase false
git clone https://github.com/ximoyuxiao/ahutoj.git
cd $WEBSERVER
git checkout -b master-test && git pull origin master-test
make build
cd core
make clean
make
judgedID=ps -ef|head -n -1|grep judged|awk '{print $2}'
kill -9 $judgedID
`./judged`