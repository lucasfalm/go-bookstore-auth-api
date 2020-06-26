package rest

import "github.com/flucas97/bookstore/users-api/utils/errors_utils"

type RestUsersRepository interface {
	Login(string, string) (*User, errors_utils.RestErr)
}

func NewRepository() RestUsersRepository {

}
