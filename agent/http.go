package agent

import (
	"net/http"
	"log"
	"time"
)

//中间件适配器

type MiddleWareFunc func(http.ResponseWriter, ...interface{}) http.Handler

//handle request method
func MethodAllow(handler http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		handler.ServeHTTP(w,r)
	})
}


// handl access log
func LoggingMiddlWare(next http.Handler) http.Handler  {

	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now(),r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

	}
	return  http.HandlerFunc(fn)


}