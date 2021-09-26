package handler

import (
	"log"
	"net/http"
	"os"
)

const (
	// VERSION version 环境变量名
	VERSION = "VERSION"
)

// New 构建 http.Handler 的实现
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)

	return logMiddleware(headerMiddleware(mux))
}

func healthz(w http.ResponseWriter, r *http.Request) {

	// 4.当访问localhost/healthz时，应返回200
	w.WriteHeader(http.StatusOK)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := wrapperResponseWriter(w)

		next.ServeHTTP(rw, r)

		// 3.Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
		log.Printf("请求 IP：%v，HTTP返回码：%v\n", r.RemoteAddr, rw.StatusCode())
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1.接收客户端request，并将request中带的header写入responseheader
		for key := range r.Header {
			w.Header().Add(key, r.Header.Get(key))
		}

		// 2.读取当前系统的环境变量中的VERSION配置，并写入responseheader
		version := os.Getenv(VERSION)
		w.Header().Add("Version", version)

		next.ServeHTTP(w, r)
	})
}
