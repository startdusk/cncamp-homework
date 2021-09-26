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
	testVersion := "test version: v0.0.1"
	os.Setenv(VERSION, testVersion)

	headers := map[string]string{
		"Accept-Encoding": "gzip",
		"User-Agent":      "Go-http-client/1.1",
		"Content-Length":  "0",
		"Date":            time.Now().Format(time.RFC3339),
		"Test-Header":     "my test header",
	}

	ts := httptest.NewServer(New())
	defer ts.Close()

	url := ts.URL + "/healthz"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	headers["Version"] = testVersion
	if len(headers) != len(res.Header) {
		t.Errorf("httpserver header 错误，响应header数量：%d，期待header数量：%d", len(res.Header), len(headers))
	}

	for k := range res.Header {
		val, ok := headers[k]
		if !ok {
			t.Errorf("httpserver header 错误，不存在 header：%v", k)
		}
		actual := res.Header.Get(k)
		if val != actual {
			t.Errorf("httpserver header 错误，header：%v，value: %v，期待value：%v", k, actual, val)
		}
	}

	version := res.Header.Get("Version")
	if version != testVersion {
		t.Errorf("httpserver VERSION 错误，得到VERSION: %s, 期待VERSION: %s", version, testVersion)
	}

	if res.StatusCode != 200 {
		t.Errorf("httpserver 响应码 错误，得到响应码: %d, 期待响应码: %d", res.StatusCode, 200)
	}
}
