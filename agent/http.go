package agent

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
)

type HttpResponse struct {
	Data interface{}  `json:"data"`
	Error error		  `json:"msg"`
}


//中间件适配器
type ServerHandlerFunc func(http.ResponseWriter, *http.Request) (int, interface{}, error)

func (fn ServerHandlerFunc) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	statusCode, data, err := fn(w, r)
	res := HttpResponse{Data:data}
	if err != nil {
		res.Error = err
	}

	jsons, err := json.Marshal(res)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	w.Write(jsons)

}

//handle request method
func MethodAllow(handler http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
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


//