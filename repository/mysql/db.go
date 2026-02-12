package mysql

import (
	"database/sql"
	"fmt"
	"main/entity"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	db *sql.DB
}

func (d *MySqlDB) IsUniquePhoneNumber(phone string) (bool, error) {
	var count int

	err := d.db.QueryRow(`SELECT COUNT(*) FROM users WHERE phone = ?`, phone).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("something went wrong: %w", err)
	}

	return count == 0, nil
}

func (d *MySqlDB) RegisterUser(u entity.UserEntity) (entity.UserEntity, error) {
	result, err := d.db.Exec(
		`INSERT INTO users(name, phone) VALUES (?, ?)`,
		u.Name,
		u.Phone,
	)

	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("can not add new user: %w", err)
	}

	id, _ := result.LastInsertId()
	u.Id = uint(id)

	return u, nil
}

func New() *MySqlDB {
	db, err := sql.Open("mysql", "gameApp:GameAppRootPass@tcp(localhost:3306)/gameApp?parseTime=true")

	if err != nil {
		panic(fmt.Errorf("connection error for db: %v", err))
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("cannot ping database: %v", err))
	}

	return &MySqlDB{db}
}
