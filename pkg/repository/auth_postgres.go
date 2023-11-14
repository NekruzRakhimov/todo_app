package repository

import (
	"fmt"
	"github.com/NekruzRakhimov/todo_app/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (id int, err error) {
	sqlQuery := fmt.Sprintf(
		`INSERT INTO %s (name, username, password_hash) VALUES($1, $2, $3) RETURNING id`, usersTable)
	if err = r.db.Table(usersTable).
		Exec(sqlQuery, user.Name, user.Username, user.Password).
		Pluck("id", &id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (u models.User, err error) {
	sqlQuery := fmt.Sprintf(
		`SELECT id, name, username FROM %s WHERE username = $1 AND password_hash = $2 `, usersTable)
	if err = r.db.Raw(sqlQuery, username, password).Scan(&u).Error; err != nil {
		return models.User{}, err
	}

	if u.ID == 0 {
		return models.User{}, gorm.ErrRecordNotFound
	}

	return u, nil
}
