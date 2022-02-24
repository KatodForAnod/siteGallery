package config

type Config struct {
	DBConfig dbConf
}

type dbConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}
