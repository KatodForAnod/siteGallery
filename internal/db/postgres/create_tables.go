package postgres

import (
	"fmt"
)

const createImgTable = `CREATE TABLE IF NOT EXISTS Img (
    Id SERIAL PRIMARY KEY,
    User_Id INTEGER REFERENCES Users (id),
    Img text,
    Tags varchar(40)[])`

//    load_by_user varchar(40) REFERENCES Users(id))`

const createUserTable = `CREATE TABLE IF NOT EXISTS Users (
    Email varchar(40) PRIMARY KEY,
    Id SERIAL UNIQUE,
    Name varchar(40),
    Password varchar(80))`

func (p *postgreSQl) CreateAllTables() error {
	if _, err := p.conn.Exec(createUserTable); err != nil {
		//log.Println("createServerDataTableExec", err)
		return fmt.Errorf("CreateAllTables createServerDataTableExec err: %s", err)
	}

	if _, err := p.conn.Exec(createImgTable); err != nil {
		//log.Println("createNodesTableExec", err)
		return fmt.Errorf("CreateAllTables createNodesTableExec err: %s", err)
	}

	return nil
}
