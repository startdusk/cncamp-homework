package handler

import (
	"fmt"
	"httpserver/metrics"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// New 构建 http.Handler 的实现
func New(lg *zap.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/hello", helloMetrics)
	// 添加指标采集
	mux.Handle("/metrics", promhttp.Handler())
	return logMiddleware(headerMiddleware(mux), lg)
}

func healthz(w http.ResponseWriter, r *http.Request) {

	// 4.当访问localhost/healthz时，应返回200
	w.WriteHeader(http.StatusOK)
}
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func helloMetrics(w http.ResponseWriter, r *http.Request) {

	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	user := r.URL.Query().Get("user")
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func logMiddleware(next http.Handler, lg *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := wrapperResponseWriter(w)

		next.ServeHTTP(rw, r)

		// 3.Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
		lg.Sugar().Infof("请求IP：%v，HTTP返回码：%v\n", ipAddress(r), rw.statusCode)
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1.接收客户端request，并将request中带的header写入response header
		for key := range r.Header {
			w.Header().Add(key, r.Header.Get(key))
		}

		// 2.读取当前系统的环境变量中的VERSION配置，并写入response header
		version := os.Getenv("VERSION")
		w.Header().Add("Version", version)

		next.ServeHTTP(w, r)
	})
}

// ipAddress 获取客户端 IP
func ipAddress(r *http.Request) string {
	if ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); ip != "" {
		return ip
	}

	if ip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
