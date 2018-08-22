package agent

import (
	"github.com/gorilla/mux"
	"github.com/asshidhak/wgo/api"
	"net/http"
)

func Register(m *mux.Router) {

	m.Handle("/", LoggingMiddlWare(MethodAllow(http.HandlerFunc(api.HelloHandle), http.MethodGet)))
}
