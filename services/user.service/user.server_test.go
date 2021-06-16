package user_service_test

import (
	"testing"
	s "../user.service"
	m "../../models"
	"time"
)

func TestCreate(t *testing.T) {

	user := m.User{
		FirstName: "John",
		LastName: "Doe",
		Email: "johndoe@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.Create(user)
	if err != nil {
		t.Error("Error trying to create a user")
		t.Fail()
	}else{
		t.Log("Success TestCreate")
	}
}

func TestRead(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}