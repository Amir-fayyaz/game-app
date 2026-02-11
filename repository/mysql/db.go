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

// // IsUniquePhoneNumber implements services.Repository.
// func (d *MySqlDB) IsUniquePhoneNumber(phone string) (bool, error) {
// 	panic("unimplemented")
// }

// // RegisterUser implements services.Repository.
// func (d *MySqlDB) RegisterUser(user entity.UserEntity) (entity.UserEntity, error) {
// 	panic("unimplemented")
// }

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
