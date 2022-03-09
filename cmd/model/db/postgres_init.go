package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"siteGallery/cmd/config"
	"siteGallery/cmd/model"
)

type postgreSQl struct {
	conn *sql.DB
}

func GetPostgresSQlConn(dbConfig config.Config) (model.Database, error) {
	log.Println("Starting initialization postgres db...")
	connStr := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		dbConfig.DBConfig.Host, dbConfig.DBConfig.Port,
		dbConfig.DBConfig.User, dbConfig.DBConfig.Password,
		dbConfig.DBConfig.Dbname, dbConfig.DBConfig.Sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	instance := postgreSQl{conn: db}
	if err = instance.CreateAllTables(); err != nil {
		log.Println(err)
		return postgreSQl{}, err
	}

	return instance, nil
}
