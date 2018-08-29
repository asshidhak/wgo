package agent

import (
	"database/sql"
	"log"
	"os"
	"time"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Config *RuntimeConfig
	Mysql  *sql.DB
	Logger *log.Logger
}

func New(config *RuntimeConfig) *Server{
	return &Server{
		Config: config,
	}
}

func (s *Server) Start()  {

	fl, _ := os.Create("wgo-" + time.Now().Format("2006-01-02") + ".log")
	s.Logger = log.New(fl, "", log.Lshortfile|log.LstdFlags)

	// bcdn:bcdn.org@tcp(192.168.3.119:3306)/BCDN?charset=utf8&parseTime=true&loc=Local
	dbstring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
						s.Config.Mysql.User,
						s.Config.Mysql.Pwd,
						s.Config.Mysql.Host,
						s.Config.Mysql.Port,
						s.Config.Mysql.Db,	)
	var dberr error
	s.Mysql, dberr = sql.Open("mysql", dbstring)
	if dberr != nil {
		s.Logger.Println(dberr)
		return
	}

	m := &mux.Router{}
	Register(m,s)
	s.Logger.Printf("Listening %s...", s.Config.Listen)
	http.ListenAndServe(s.Config.Listen , m)


}
