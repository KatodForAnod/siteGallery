package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	DBConfig dbConf `json:"db_config"`
	SvConfig server `json:"sv_config"`
}

type dbConf struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
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
		return Config{}, fmt.Errorf("load config err: %s", err)
	}

	err = json.Unmarshal(data, &loadedConf)
	return
}
