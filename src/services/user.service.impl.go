package services

import (
	"log"
	"github.com/keithyw/kyw-go-docker-test/grpc"
	"github.com/keithyw/kyw-go-docker-test/models"
	"github.com/keithyw/kyw-go-docker-test/repositories"
	"github.com/keithyw/kyw-go-docker-test/utils"
)

type UserServiceImpl struct {
	grpcClient *grpc.Client
	repo repositories.UserRepository
}

func NewUserService (client *grpc.Client, repo repositories.UserRepository) UserService {
	return &UserServiceImpl{client, repo}
}

func (us *UserServiceImpl) CreateUser(user models.User) (*models.User, error) {
	if len(user.Passwd) > 0  {
		passwd, err := utils.Encrypt(user.Passwd)
		if err != nil {
			return nil, err
		}
		user.Passwd = passwd
	}
	newUser, err := us.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	us.grpcClient.CreateUser(newUser)
	return newUser, nil
}

func (us *UserServiceImpl) DeleteUser(id int) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	us.grpcClient.DeleteUser(id)
	return nil
}

func (us *UserServiceImpl) UpdateUser(id int, user models.User) (*models.User, error) {
	if len(user.Passwd) > 0  {
		passwd, err := utils.Encrypt(user.Passwd)
		if err != nil {
			return nil, err
		}
		user.Passwd = passwd
	}
	updatedUser, err := us.repo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	us.grpcClient.UpdateUser(id, updatedUser)
	return updatedUser, nil
}

func (us *UserServiceImpl) FindUserById(id int) (*models.User, error) {
	user, err := us.repo.FindUserById(id)
	if err != nil {
		log.Println("FindByUserId failed: " + err.Error())
		return nil, err
	}
	return user, err
}

func (us *UserServiceImpl) FindUserByName(name string) (*models.User, error) {
	user, err := us.repo.FindUserByName(name)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (us *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := us.repo.GetAllUsers()
	if err != nil {
		log.Printf("GetAllUsers error %s", err.Error())
		return nil, err
	}
	return users, nil
}