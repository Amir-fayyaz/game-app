package mysql

import (
	"database/sql"
	"fmt"
	"main/entity"
)

func (d *MySqlDB) IsPhoneNumber(phone string) (bool, error) {

	user := &entity.UserEntity{}

	row := d.db.QueryRow(`select * from users where phone=?`, phone)

	err := row.Scan(&user.Id, &user.Name, &user.Phone, new(interface{}))

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("something went wrong : %w", err)
	}

	return false, nil

}

func (d *MySqlDB) Register(u entity.UserEntity) (entity.UserEntity, error) {
	result, err := d.db.Exec(`insert into users(name,phone) values (?,?)`, u.Name, u.Phone)

	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("can not add new user, %w", err)
	}

	id, _ := result.LastInsertId()

	u.Id = uint(id)

	return u, nil
}
