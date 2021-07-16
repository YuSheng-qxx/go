package conf

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var Sys *Sysconfig = &Sysconfig{}

func init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, Sys)
	if err != nil {
		panic(err)
	}
}

type Sysconfig struct {
	Port         string    `json:"Port"`
	DBUserName   string    `json:"DBUserName"`
	DBPassword   string    `json:"DBPassword"`
	DBIp         string    `json:"DBIp"`
	DBPort       string    `json:"DBPort"`
	DBName 		 string    `json:"DBName"`
	Collection   string    `json:"Collection"`
}