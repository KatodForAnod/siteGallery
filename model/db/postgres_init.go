package db

import (
	"database/sql"
	"fmt"
	"log"
	"siteGallery/config"
	"siteGallery/model"
)

type postgreSQl struct {
	conn *sql.DB
}

func GetPostgresSQlConn(dbConfig config.Config) (model.Database, error) {
	log.Println("Starting initialization postgres db...")
	connStr := fmt.Sprintf(`user=%s password=%s dbname=%s sslmode=%s`,
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

	return postgreSQl{conn: db}, nil
}
