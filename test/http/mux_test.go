package http

import (
	"fmt"
	"net/http"
	"testing"
)

// 处理http请求处理 要实现http.Handler接口
type MsgHandler struct {
	msg string
}

func (m MsgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(m.msg))
}

// http.NewServeMux创建路由容器
// http.ServeMux#Handle 注入路由和http.Handler到mux里的map
// http.ServeMux也是一个http.Handler 它最终传入http.ListenAndServe处理请求
func TestMsgHandler(t *testing.T) {
	mux := http.NewServeMux()

	mux.Handle("/", MsgHandler{"hello world"})
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	http.ListenAndServe(":1234", mux)
}

// middleware
func m1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("m1 start")
		next.ServeHTTP(w, r)
		fmt.Println("m1 end")
	})
}

func m2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("m2 start")
		next.ServeHTTP(w, r)
		fmt.Println("m2 end")
	})
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", m1(m2(MsgHandler{"hello middleware"})))
	http.ListenAndServe(":1234", mux)
}

func TestFileServer(t *testing.T) {
	http.ListenAndServe(":1234", http.FileServer(http.Dir("/tmp")))
}
