package access_token

import (
	"github.com/flucas97/bookstore/auth-api/src/repository/db"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

type Service interface {
	GetById(string) (*AccessToken, error)
}

type service struct {
	dbRepo db.DbRepository
}

func NewService(dbRepo db.DbRepository) *service {
	return &service{
		dbRepo: dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.RestErr) {
	return nil, nil
}
