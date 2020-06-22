package access_token

import "github.com/flucas97/bookstore/users-api/utils/errors_utils"

type Repository interface {
	GetById(string) (*accessTokenService, *errors_utils.RestErr)
}

type Service interface {
	GetById(string) (*accessTokenService, *errors_utils.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service { 
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(ID string) (*accessTokenService, *errors_utils.RestErr)
	return nil, nil 
}