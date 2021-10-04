# 启动编译环境
FROM golang:1.15-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中
COPY . /go/src/httpserver

# 编译
WORKDIR /go/src/httpserver
RUN go install .

FROM alpine:3.13

COPY --from=builder /go/bin/httpserver /bin/httpserver

ENV ADDR=:80

# 申明暴露端口
EXPOSE 80

# 设置服务入口
ENTRYPOINT [ "/bin/httpserver" ]
