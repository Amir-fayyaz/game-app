package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	db *sql.DB
}

func New() *MySqlDB {
	db, err := sql.Open("mysql", "gameApp:GameAppRootPass@(localhost:3306)/gameApp")

	if err != nil {
		panic(fmt.Errorf("connection error for db : %v", err))
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{db}
}
