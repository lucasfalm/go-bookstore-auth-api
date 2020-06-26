package rest

import (
	"github.com/flucas97/bookstore/users-api/model/users"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

type RestUsersRepository interface {
	Login(string, string) (*users.User, errors_utils.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) Login(email string, password string) (*users.User, errors_utils.RestErr) {

}
