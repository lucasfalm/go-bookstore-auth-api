package db

import (
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

// entrypoint to use the repository
type DbRepositoryInterface interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

func NewRepository() DbRepositoryInterface {
	return &dbRepository{}
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors_utils.RestErr) {
	return nil, nil
}
