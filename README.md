# 构建项目

运行 `env GOOS=linux go build .` 编译 Linux 系统中可以执行的二进制文件。
构建成功后需要赋予二进制文件可以执行的权限 `chmod +x main`。

# Build

Please run `env GOOS=linux go build .` for builiding linux runable binary.
And you should make sure the binary has execute permission `chmod +x main`.

# 交叉编译

```shell
  GOARCH=amd64 GOOS=linux go build -ldflags "-s -w"
```

# 安装FUN脚手架

```shell
  sudo npm install @alicloud/fun -g
```