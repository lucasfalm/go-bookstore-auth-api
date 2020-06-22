package access_token

import (
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
}
type ServiceInterface interface {
	GetById(string) (*AccessToken, *errors_utils.RestErr)
}

type service struct {
	dbRepo Repository
}

func NewService(repo Repository) ServiceInterface {
	return &service{
		dbRepo: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.RestErr) {
	return nil, nil
}
