package repositories

import "github.com/keithyw/kyw-go-docker-test/models"

type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, user models.User) (*models.User, error)

	FindUserById(id int) (*models.User, error)
	FindUserByName(name string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}