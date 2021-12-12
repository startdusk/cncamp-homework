package main

import (
	"context"
	"httpserver/handler"
	"httpserver/logger"
	"httpserver/metrics"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namsral/flag"
)

// 优先读取k8s设置得环境变量
var addr = flag.String("addr", ":80", "address to listen")
var version = flag.String("version", "v0.0.1", "version for httpserver")
var shotdownTime = flag.Int("shotdown_time", 10, "shotdownTime for httpserver")
var logFile = flag.String("log_file", "httpserver.log", "log_file for httpserver")
var logLevel = flag.String("log_level", "INFO", "log_level for httpserver")

func main() {
	flag.Parse()

	os.Setenv("VERSION", *version)

	lg, err := logger.New(*logFile, *logLevel)
	if err != nil {
		log.Fatalf("cannot create a logger: %v", err)
	}
	defer lg.Sync()
	sugar := lg.Sugar()

	if err := metrics.Register(); err != nil {
		log.Fatalf("prometheus register error: %v", err)
	}

	h := handler.New(lg)
	srv := http.Server{
		Addr:    *addr,
		Handler: h,
	}

	// 下面 graceful-shutdown 参考: https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalf("listen: %s\n", err)
		}
	}()

	sugar.Info("httpserver start on", *addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sugar.Info("httpserver shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*shotdownTime)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatal("httpserver forced to shutdown: ", err)
	}

	sugar.Info("httpserver exiting")
}
