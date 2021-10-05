# cncamp-homework

## 1.httpserver

### 目录：/cncamp-homework/httpserver

Docker 运行:

```bash
$ make run-image
```

二进制运行 (需要管理员权限)：

```bash
$ make run
```

httpserver 启动后，访问 http://localhost/healthz

当然是看不到什么的，要用单元测试验证代码的正确性，[这里](https://github.com/startdusk/cncamp-homework/blob/master/httpserver/handler/handler_test.go)
