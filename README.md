

#### 安装教程

1.  切换 GO 镜像

    go env -w GOPROXY=https://goproxy.cn,direct

    清除缓存 go clean -modcache

    验证镜像切换 go env | grep GOPROXY

    找到 GOPROXY=https://goproxy.cn,direct

2.  go mod tidy 下载依赖
3.  go run ./main.go  启动项目
4.  go build -o serve ./main.go 打包项目

#### 使用说明


1.  cmd/bs.sh 文件快捷生成，在 handler、model、router、service 这4个目录生成 go 文件，Windows 用 Git Bash 执行,【注意小写字母】
2.  pkg/tables 是数据库自动迁移的的配置类，注意数据库注释
3.  pkg/serr是业务错误码，用于隐藏错误信息
4.  config是配置的类
5.  router是路由地址，在internal中的fiber实例都需要router的实例。
