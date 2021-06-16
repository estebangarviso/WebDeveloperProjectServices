package user_service

import (
	m "../../models"
	repo "../../repositories/user.repository"
)

func Create(user m.User) error {

	err := repo.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {

	users, err := repo.Read()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func Update(user m.User, userID string) error {

	err := repo.Update(user, userID)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userID string) error {

	err := repo.Delete(userID)

	if err != nil {
		return err
	}

	return nil
}