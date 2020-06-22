package db

import (
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

type DatabaseRepository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

type databaseRepository struct{}

func NewRepository() *databaseRepository {
	return &databaseRepository{}
}

func (r *databaseRepository) GetById(string) (*access_token.AccessToken, *errors_utils.RestErr) {
	return nil, nil
}
