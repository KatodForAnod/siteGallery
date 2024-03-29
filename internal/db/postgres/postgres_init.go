package postgres

import (
	"KatodForAnod/siteGallery/internal/config"
	"KatodForAnod/siteGallery/internal/db"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type postgreSQl struct {
	conn *sql.DB
}

func GetPostgresSQlConn(dbConfig config.Config) (db.Database, error) {
	log.Println("Starting initialization postgres db...")
	connStr := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		dbConfig.DBConfig.Host, dbConfig.DBConfig.Port,
		dbConfig.DBConfig.User, dbConfig.DBConfig.Password,
		dbConfig.DBConfig.Dbname, dbConfig.DBConfig.Sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("GetPostgresSQlConn err:%s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("GetPostgresSQlConn err:%s", err)
	}

	instance := postgreSQl{conn: db}
	if err = instance.CreateAllTables(); err != nil {
		return postgreSQl{}, fmt.Errorf("GetPostgresSQlConn err:%s", err)
	}

	return instance, nil
}
