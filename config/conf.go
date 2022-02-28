package config

type Config struct {
	DBConfig dbConf
	SvConfig server
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
