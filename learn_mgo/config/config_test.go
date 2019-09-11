package config

import (
	"encoding/json"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}

func TestLoadConfig(t *testing.T) {
	t.Log("start test LoadConfig")
	var conf MgoConfig
	_ = LoadConfig("mgo.json", &conf)
	s, _ := json.MarshalIndent(conf, "\t", "")
	log.Print(string(s))
}
