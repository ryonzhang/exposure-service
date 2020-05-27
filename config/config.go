package config

import (
	"github.com/juvoinc/exposure-service/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Instance *model.Config

func init() {
	Instance = &model.Config{}
	yamlFile, err := ioutil.ReadFile("metrics.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, Instance)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return
}
