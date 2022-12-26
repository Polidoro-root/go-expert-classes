package database

import "github.com/Polidoro-root/go-expert/7_apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
