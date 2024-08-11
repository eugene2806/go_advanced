package storage

import (
	"fmt"
	"go_advanced/Lesson7/internal/app/models"
	"log"
)

type UserRepository struct {
	storage *Storage
}

var tableUser = "users"

//Добавить юзера в БД

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)

	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {

		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool

	if err != nil {
		return nil, false, err
	}

	var userFind *models.User

	for _, u := range users {
		if u.Login == login {
			userFind = u
			founded = true

			break
		}
	}

	return userFind, founded, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT* FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	users := make([]*models.User, 0, 20)

	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.Login, &u.Password)

		if err != nil {
			log.Println(err)
			continue
		}

		users = append(users, &u)
	}

	return users, nil
}
