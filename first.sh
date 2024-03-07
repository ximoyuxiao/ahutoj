#安装docker环境
if command -v docker >/dev/null 2>&1; then
    echo "Docker is installed"
else
    echo "Docker is not installed, start installing..."
    bash <(curl -sSL https://gitee.com/SuperManito/LinuxMirrors/raw/main/DockerInstallation.sh)
fi

#创建日志文件,挂载到容器中

sudo mkdir -p .logs/origin && sudo touch "${DIR:-./tmp}"/origin/log/ahutoj.log
sudo mkdir -p "${DIR:-./tmp}"/gateway/log&& sudo touch "${DIR:-./tmp}"/gateway/log/ahutoj.log
sudo mkdir -p "${DIR:-./tmp}"/persistence/log &&sudo  touch "${DIR:-./tmp}"/persistence/log/ahutoj.log
sudo mkdir -p "${DIR:-./tmp}"/oj/log && sudo touch "${DIR:-./tmp}"/oj/log/ahutoj.log
sudo chmod -R 777 "${DIR:-./tmp}"/

#运行容器和删除构建中间镜像
sudo docker compose up -d

#修复npm容器zope环境
docker exec -it oj-npm bash -c "python3 -m pip install --upgrade pip &&
sed -i 's#dl-cdn.alpinelinux.org#mirrors.aliyun.com#g' /etc/apk/repositories &&
apk add  build-base &&
pip uninstall  cffi -y &&
apk add python3-dev &&
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple cffi certbot-dns-dnspod zope &&
exit"

# shellcheck disable=SC2046
sudo docker rmi $(sudo docker images --filter "dangling=true" -q)