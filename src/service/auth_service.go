package service

import (
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/repository/db"
	"github.com/flucas97/bookstore/auth-api/src/repository/rest"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

// related with controller
type ServiceInterface interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
	Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors_utils.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors_utils.RestErr
}

type service struct {
	dbRepo    db.DbRepositoryInterface
	usersRepo rest.RestUsersRepository
}

// wich repository the service will use
func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepositoryInterface) ServiceInterface {
	return &service{
		dbRepo:    dbRepo,
		usersRepo: usersRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors_utils.RestErr) {
	AccessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return AccessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors_utils.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.usersRepo.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}
