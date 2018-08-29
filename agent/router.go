package agent

import (
	"github.com/gorilla/mux"
	"github.com/asshidhak/wgo/api"
	"net/http"
)

func Register(m *mux.Router, s *Server) {

	m.Handle("/", MethodAllow(http.HandlerFunc(api.HelloHandle), http.MethodGet))
	m.Handle("/login", MethodAllow(ServerHandlerFunc(s.HandleUserLogin), http.MethodGet))
	m.Handle("/login", MethodAllow(ServerHandlerFunc(s.HandleUserLogin), http.MethodGet))
}
