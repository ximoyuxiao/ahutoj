#安装docker环境
if command -v docker >/dev/null 2>&1; then
    echo "Docker is installed"
else
    echo "Docker is not installed, start installing..."
    bash <(curl -sSL https://gitee.com/SuperManito/LinuxMirrors/raw/main/DockerInstallation.sh)
fi

#创建日志文件,挂载到容器中
sudo mkdir -p "${DIR:-.}"/origin/log && sudo touch "${DIR:-.}"/origin/log/ahutoj.log
sudo mkdir -p "${DIR:-.}"/gateway/log&& sudo touch "${DIR:-.}"/gateway/log/ahutoj.log
sudo mkdir -p "${DIR:-.}"/persistence/log &&sudo  touch "${DIR:-.}"/persistence/log/ahutoj.log
sudo mkdir -p "${DIR:-.}"/oj/log && sudo touch "${DIR:-.}"/oj/log/ahutoj.log

#运行容器和删除构建中间镜像
sudo docker compose up -d
# shellcheck disable=SC2046
sudo docker rmi $(sudo docker images --filter "dangling=true" -q)

#修复npm容器zope环境
docker exec -it oj-npm bash
python3 -m pip install --upgrade pip
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
apk add  build-base
pip uninstall  cffi
pip install  cffi
apk add python3-dev
pip install certbot-dns-dnspod
pip install zope
exit