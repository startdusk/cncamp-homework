package main

import (
	"context"
	"flag"
	"httpserver/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", ":8080", "address to listen")
var version = flag.String("version", "v0.0.1", "version for httpserver")

func main() {
	flag.Parse()

	os.Setenv("VERSION", *version)

	h := handler.New()
	srv := http.Server{
		Addr:    *addr,
		Handler: h,
	}

	// 下面 graceful-shutdown 参考: https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("httpserver start on", *addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("httpserver shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("httpserver forced to shutdown: ", err)
	}

	log.Println("httpserver exiting")
}
