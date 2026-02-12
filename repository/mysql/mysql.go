package mysql

import (
	"database/sql"
	"fmt"
	"main/entity"
)

func (d *MySqlDB) IsPhoneNumberExists(phone string) (bool, error) {
	var count int
	err := d.db.QueryRow(`SELECT COUNT(*) FROM users WHERE phone = ?`, phone).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking phone: %w", err)
	}
	return count > 0, nil
}

func (d *MySqlDB) GetUserByPhone(phone string) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}

	err := d.db.QueryRow(
		`SELECT id, name, phone FROM users WHERE phone = ?`,
		phone,
	).Scan(&user.Id, &user.Name, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (d *MySqlDB) GetUserByID(id uint) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}

	err := d.db.QueryRow(
		`SELECT id, name, phone FROM users WHERE id = ?`,
		id,
	).Scan(&user.Id, &user.Name, &user.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (d *MySqlDB) GetAllUsers() ([]entity.UserEntity, error) {
	rows, err := d.db.Query(`SELECT id, name, phone FROM users ORDER BY id DESC`)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %w", err)
	}
	defer rows.Close()

	var users []entity.UserEntity
	for rows.Next() {
		var user entity.UserEntity
		err := rows.Scan(&user.Id, &user.Name, &user.Phone)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *MySqlDB) UpdateUser(user entity.UserEntity) error {
	_, err := d.db.Exec(
		`UPDATE users SET name = ?, phone = ? WHERE id = ?`,
		user.Name,
		user.Phone,
		user.Id,
	)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (d *MySqlDB) DeleteUser(id uint) error {
	_, err := d.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
