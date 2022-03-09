package db

import "log"

const createImgTable = `CREATE TABLE IF NOT EXISTS Img (
    id SERIAL PRIMARY KEY,
    img text,
    tags varchar(40)[])`

//    load_by_user varchar(40) REFERENCES Users(id))`

const createUserTable = `CREATE TABLE IF NOT EXISTS Users (
    email varchar(40) PRIMARY KEY,
    id SERIAL,
    name varchar(40) )`

func (p *postgreSQl) CreateAllTables() error {
	if _, err := p.conn.Exec(createUserTable); err != nil {
		log.Println("createServerDataTableExec", err)
		return err
	}

	if _, err := p.conn.Exec(createImgTable); err != nil {
		log.Println("createNodesTableExec", err)
		return err
	}

	return nil
}
