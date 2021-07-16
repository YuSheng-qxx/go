package datasource

import (
	"app/test/conf"
	"gopkg.in/mgo.v2"
	"log"
)

var c *mgo.Collection

func GetCollection() *mgo.Collection {
	return c
}

func init() {
	path := conf.Sys.DBIp + ":" + conf.Sys.DBPort
	sion, err := mgo.Dial(path)
	db:= sion.DB(conf.Sys.DBName)
	c = db.C(conf.Sys.Collection)

	if err !=nil{
		log.Println(err)
	}
}