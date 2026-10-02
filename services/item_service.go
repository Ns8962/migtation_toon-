package services

import "myapi/models"

type UserService struct {
	users []models.User
}

// constructor
func NewUserService() *UserService {
	return &UserService{
		users: []models.User{},
	}
}

// method (เหมือน OOP)
func (s *UserService) CreateUser(user models.User) models.User {
	s.users = append(s.users, user)
	return user
}

func (s *UserService) GetUsers() []models.User {
	return s.users
}
