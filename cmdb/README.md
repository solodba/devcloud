# CMDB(配置管理)微服务

## docker安装MySQL
```sh
docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql mariadb:latest
```