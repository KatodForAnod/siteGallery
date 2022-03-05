package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	DBConfig dbConf `json:"db_config"`
	SvConfig server `json:"sv_config"`
}

type dbConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

type server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

const configPath = "conf.config"

func LoadConfig() (loadedConf Config, err error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

	err = json.Unmarshal(data, &loadedConf)
	return
}
