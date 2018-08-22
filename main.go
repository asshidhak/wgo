package main

import (
	"github.com/asshidhak/wgo/agent"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {



	log.SetFlags(log.Lshortfile | log.LstdFlags)


	cfg, err := agent.LoadConfig("app.json")
	if err != nil {
		log.Printf("Load cfg file faild: %v", err)
		return
	}
	a := agent.New(cfg)
	a.Start()


}
