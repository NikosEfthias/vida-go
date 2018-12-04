package user

import (
	"fmt"

	"gitlab.mugsoft.io/vida/go-api/models"
)

func Service_register(name, lastname, email, phone, password string) (string, error) {
	switch {
	case "" == email && "" == phone:
		return "", fmt.Errorf("missing email and phone")
	case "" == name:
		return "", fmt.Errorf("missing name")
	case "" == lastname:
		return "", fmt.Errorf("missing lastname")
	case "" == password:
		return "", fmt.Errorf("missing password")
	}
	err := models.User_new(&models.User{
		Name:     name,
		Lastname: lastname,
		Password: password,
		Email:    email,
		Phone:    phone,
	})
	if nil != err {
		return "", err
	} else {
		return "success", nil
	}
}
