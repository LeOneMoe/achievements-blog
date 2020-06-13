package db

import "database/sql"

type DataBase struct {
	Server     string `json:"server"`
	Port       string `json:"port"`
	Driver     string `json:"driver"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Connection *sql.DB
}

type Connector interface {
	GetConnection()
	CloseConnectionIfExists()
}

func (db *DataBase) GetConnection() error {
	conn, err := sql.Open(db.Driver+":"+db.Port, db.Server)
	if err != nil {
		return err
	}

	db.Connection = conn
	return nil
}

func (db *DataBase) CloseConnectionIfExists() error {
	if err := db.Connection.Close(); err != nil {
		return err
	}

	return nil
}

//var bd Connector = &DataBase {
//	URL:    "123",
//	driver: "123",
//	User:   User {
//		user: "123",
//		pass: "123",
//	},
//}
