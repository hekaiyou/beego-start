# beego-start

## 运行环境

安装 [GO](https://golang.google.cn/) 语言环境，并配置好下面两个环境变量。

- GOROOT: Go的安装目录
- GOPATH: Go项目代码和第三方依赖包目录（可选）
  - Linux 默认值：$HOME/go
  - Windows 默认值：%USERPROFILE%\go

### Beego框架

安装 [Beego](https://beego.me/) Web 应用开发框架。

```shell
go get github.com/beego/beego/v2@v2.0.1
go get github.com/beego/beego/v2/core/config
go get github.com/beego/bee
```

将 `$GOPATH/bin` 加入到你 `PATH` 变量中，确保 `Bee` 命令可以正常使用。

```shell
vim ~/.bashrc
------
export PATH=$PATH:/root/go/bin
------
source ~/.bashrc
```

### 数据存储

连接 **MongoDB** 数据库，项目会读取 `MONGODB_HOST` 和 `MONGODB_PORT` 两个环境变量，读取不到时默认使用 *127.0.0.1:27017* 连接数据库。

```shell
go get github.com/globalsign/mgo
```

### 运行项目

初始化模块依赖。

```shell
go mod init beego_start
go get -d -v ./...
```

热编译运行项目。

```shell
bee run
```

运行项目并自动生成 **Swagger** API 文档，通过访问 *http://127.0.0.1:5000/swagger/* 路径打开文档页面。

```shell
bee run -gendoc=true -downdoc=true
```

*注意：生成 **Swagger** 文档有时会有缓存，需要删除浏览器缓存及项目中的文档，重新生成即可。*
