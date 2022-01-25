package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	REDIS_HOST     string
	REDIS_PORT     string
	STORAGE_URL    string
	STORAGE_ID     string
	STORAGE_SECRET string
	CONFIG_SMTP_HOST       string
	CONFIG_SMTP_PORT       int
	CONFIG_SMTP_AUTH_EMAIL string
	CONFIG_AUTH_PASSWORD   string
	CONFIG_SENDER_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("./config.json", &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

func GetConfigTest() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("./../../config.json", &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
