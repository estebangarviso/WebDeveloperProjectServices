package user_repository

import (
	m "../../models"
)

func Create(user m.User) error {
	return nil
}

func Read() (m.Users, error) {

	var users m.Users

	return users, nil
}

func Update(user m.User, userID string) error {
	return nil
}

func Delete(userID string) error {
	return nil
}