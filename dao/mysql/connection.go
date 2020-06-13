package mysql

import (
	"database/sql"
	"fmt"
	"github.com/LeOneMoe/go-gin-react-crud/utilities"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getConnection() *sql.DB {
	cfg, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Pass, cfg.Server, cfg.Port, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
