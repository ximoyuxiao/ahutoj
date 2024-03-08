#!/bin/bash

# 设置数据库连接参数
DB_HOST="localhost"
DB_PORT="3306"
DB_USER="root"
DB_PASSWORD=$MYSQL_ROOT_PASSWORD
DB_NAME="ahutoj"

# 设置备份目录和文件名
BACKUP_DIR="/backup"
BACKUP_FILE="$BACKUP_DIR/ahutoj_backup_$(date +%Y%m%d%H%M%S).sql"

# 备份数据库
mysqldump -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASSWORD --single-transaction --databases $DB_NAME > $BACKUP_FILE

# 输出备份完成信息
echo "备份已完成，文件保存在：$BACKUP_FILE"

file_count=$(ls -1 $BACKUP_DIR/*.sql 2>/dev/null | wc -l)
if [[ $file_count -gt 3 ]]; then
  echo "备份文件数量超过3个，执行删除操作"
  oldest_file=$(ls -t $BACKUP_DIR/*.sql | tail -n 1)
  rm $oldest_file
fi
apt update

apt list --installed|grep sshpass
if [ $? -ne 0 ]; then
  echo "正在安装sshpass..."
  apt install -y sshpass
fi

apt list --installed|grep cron
if [ $? -ne 0 ]; then
  echo "正在安装cron..."
  apt install -y cron
fi

ssh-keygen -R $REMOTE_BACKUP_SERVER
# 只要有一个为空就不进行
if [ -z $REMOTE_BACKUP_SERVER ] || [ -z $REMOTE_BACKUP_USER ] || [ -z $REMOTE_BACKUP_PASSWORD ] || [ -z $REMOTE_BACKUP_PATH ]; then
  echo "远程备份参数不完整，不进行远程备份"
  # 打印所有参数
  echo "REMOTE_BACKUP_SERVER: $REMOTE_BACKUP_SERVER"
  echo "REMOTE_BACKUP_USER: $REMOTE_BACKUP_USER"
  echo "REMOTE_BACKUP_PASSWORD: $REMOTE_BACKUP_PASSWORD"
  echo "REMOTE_BACKUP_PATH: $REMOTE_BACKUP_PATH"
  exit 0
else 
  sshpass -p $REMOTE_BACKUP_PASSWORD scp -o StrictHostKeyChecking=no $BACKUP_FILE $REMOTE_BACKUP_USER@$REMOTE_BACKUP_SERVER:$REMOTE_BACKUP_PATH
  if [ $? -eq 0 ]; then
    echo "远程备份已完成，文件保存在：$REMOTE_BACKUP_SERVER:$REMOTE_BACKUP_PATH"
  else
    echo "远程备份失败"
  fi
fi

cat /var/spool/cron/crontabs/root|grep "/backup.sh"
if [ $? -ne 0 ]; then
  echo "正在添加定时任务..."
  touch /var/spool/cron/crontabs/root
  echo "0 3 * * * /backup.sh" >> /var/spool/cron/crontabs/root
fi
