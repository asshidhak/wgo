package agent

import (
	"io/ioutil"
	"encoding/json"
)

type RuntimeConfig struct {
	Listen string   `json:"listen"`
	Mysql struct{
		Host string `json:"host"`
		Port int    `json:"port"`
		User string `json:"user"`
		Pwd  string `json:"pwd"`
		Db   string `json:"db"`
	}
}


func LoadConfig(confpath string) (*RuntimeConfig, error) {

	conf := &RuntimeConfig{}
	bytes , err := ioutil.ReadFile(confpath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil

}
