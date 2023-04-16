package services

import (
	"github.com/Islam-Miko/go-mongodb/models"
)


type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}