package repository

import (
	"gorm-tutorial/src/modules/domain"
)

// output
type Output struct {
	Result interface{}
	Error  error
}

type UserRepository interface {
	Save(*domain.User) Output
	Update(*domain.User) Output
	Delete(*domain.User) Output
	FindAll() Output
	FindByID(string) Output
}
