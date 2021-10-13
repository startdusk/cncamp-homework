# httpserver

- [httpserver](#httpserver)
  - [1.作业一](#1作业一)
  - [2.作业二](#2作业二)

## 1.作业一

在当前目录下

Docker 运行:

```bash
$ make run-image
```

二进制运行 (需要管理员权限，适用于 Linux，MacOS，Windows10 WSL2 上运行)：

```bash
$ make run
```

httpserver 启动后，访问 http://localhost/healthz

当然是看不到什么的，要用单元测试验证代码的正确性，[这里](https://github.com/startdusk/cncamp-homework/blob/master/httpserver/handler/handler_test.go)

## 2.作业二

Docker 运行:

```bash
$ make run-image
```

推送到 DockerHub:

```bash
$ make upload-image
```
