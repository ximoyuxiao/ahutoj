#安装docker环境
if command -v docker >/dev/null 2>&1; then
    echo "Docker is installed"
else
    echo "Docker is not installed, start installing..."
    bash <(curl -sSL https://gitee.com/SuperManito/LinuxMirrors/raw/main/DockerInstallation.sh)
fi

#创建日志文件,挂载到容器中

sudo mkdir -p ./logs/origin && sudo touch ./logs/origin/ahutoj.log
sudo mkdir -p ./logs/gateway&& sudo touch ./logs/gateway/ahutoj.log
sudo mkdir -p ./logs/persistence &&sudo  touch ./logs/persistence/ahutoj.log
sudo mkdir -p ./logs/oj && sudo touch ./logs/oj/ahutoj.log
sudo chmod -R 777 ./logs

#运行容器和删除构建中间镜像
sudo docker compose up -d
sudo docker compose -f env.docker-compose.yml up -d

#修复npm容器zope环境
# docker exec -it oj-npm bash -c "python3 -m pip install --upgrade pip &&
# sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories &&
# apk add  build-base &&
# pip uninstall  cffi -y &&
# apk add python3-dev &&
# pip install -i https://pypi.tuna.tsinghua.edu.cn/simple cffi certbot-dns-dnspod zope &&
# exit"

docker exec -it oj-mysql bash -c "bash /backup.sh && exit"

# shellcheck disable=SC2046
sudo docker rmi $(sudo docker images --filter "dangling=true" -q)

