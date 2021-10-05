package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// 参考: https://pkg.go.dev/net/http/httptest
func TestHandler(t *testing.T) {
	// 测试数据：环境变量VERSION和预设请求头部的header
	testVersion := "test version: v0.0.1"
	os.Setenv("VERSION", testVersion)

	headers := map[string]string{
		"Accept-Encoding": "gzip",
		"User-Agent":      "Go-http-client/1.1",
		"Content-Length":  "0",
		"Date":            time.Now().Format(time.RFC3339),
		"Test-Header":     "my test header",
	}

	// 启动我们的httpserver，并构造get 请求(带上预设的header) 发起请求 /healthz
	ts := httptest.NewServer(New())
	defer ts.Close()

	url := ts.URL + "/healthz"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 设置header到请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// 返回的header数量必须与预设的header+VERSION一致
	headers["Version"] = testVersion
	if len(headers) != len(res.Header) {
		t.Errorf("httpserver header 错误，响应header数量：%d，期望header数量：%d", len(res.Header), len(headers))
	}

	for k := range res.Header {
		val, ok := headers[k]
		if !ok {
			t.Errorf("httpserver header 错误，不存在 header：%v", k)
		}
		actual := res.Header.Get(k)
		if val != actual {
			t.Errorf("httpserver header 错误，header：%v，value: %v，期望value：%v", k, actual, val)
		}
	}

	// 检查header中返回的VERSION环境变量
	version := res.Header.Get("Version")
	if version != testVersion {
		t.Errorf("httpserver VERSION 错误，得到VERSION: %s, 期望VERSION: %s", version, testVersion)
	}

	// 检查访问/healthz返回状态码 == 200
	if res.StatusCode != 200 {
		t.Errorf("httpserver 响应码 错误，得到响应码: %d, 期望响应码: %d", res.StatusCode, 200)
	}
}
