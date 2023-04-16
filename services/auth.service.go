package services

import (
	"github.com/Islam-Miko/go-mongodb/models"
)


type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}