package main

import (
	"encoding/json"
	"fmt"
	"learn_mgo/config"
	"learn_mgo/entity"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

func main() {
	var conf config.MgoConfig
	_ = config.LoadConfig("learn_mgo/config/mgo.json", &conf)
	Show(conf)
	dialInfo := mgo.DialInfo{
		Addrs:    []string{conf.Host},
		Database: conf.Database,
	}
	fmt.Println("----------")
	fmt.Println(dialInfo)
	session, err := mgo.DialWithInfo(&dialInfo)
	fmt.Println("err: ", err)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	col := session.DB(conf.Database).C(conf.Collection)
	err = col.Insert(bson.M{"name": "godme", "age": 23, "gender": "male"})
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------")
	results := make([]entity.Person, 0)
	_ = col.Find(bson.M{}).All(&results)
	for _, result := range results {
		Show(result)
	}
}

func Show(source interface{}) {
	bs, _ := json.MarshalIndent(source, "\t", "")
	fmt.Println(string(bs))
}
