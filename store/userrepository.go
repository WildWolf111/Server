package store

import (
	"StandartWebServer/internal/models"
	"fmt"
	"log"
)

type UserRepository struct {
	store *Store
}

var (
	tableUser string = "users"
)

func (ur *UserRepository) Created(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s(login, password) VALUES($1, $2) RETURNING ID	", tableUser)
	if err := ur.store.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}

		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return nil, nil
}
