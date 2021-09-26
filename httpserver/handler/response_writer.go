package handler

import "net/http"

// responseWriter 重写 http.ResponseWriter，
// 原生 http.ResponseWriter 会被拦截，不会返回给用户处理，因此我们也拿不到 http 返回的响应码
// 所以我们包一下http.ResponseWriter，让它带点信息
type responseWriter struct {
	w http.ResponseWriter

	// statusCode 保存http响应码
	statusCode int
}

// Header 实现http.ResponseWriter接口
func (r *responseWriter) Header() http.Header {
	return r.w.Header()
}

// Write 实现http.ResponseWriter接口
func (r *responseWriter) Write(bs []byte) (int, error) {
	return r.w.Write(bs)
}

// WriteHeader 实现http.ResponseWriter接口
func (r *responseWriter) WriteHeader(statusCode int) {
	// 此处拦截http的返回码
	r.statusCode = statusCode
	r.w.WriteHeader(statusCode)
}

// StatusCode 返回http响应码
func (r *responseWriter) StatusCode() int {
	return r.statusCode
}

func wrapperResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		w: w,
	}
}
