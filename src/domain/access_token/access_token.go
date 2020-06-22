package access_token

import "github.com/flucas97/bookstore/users-api/utils/errors_utils"

var (
	AccessTokenService AccessTokenServiceInterface = &accessTokenService{}
)

type accessTokenService struct {
	repository Repository
}

type AccessTokenServiceInterface interface {
	GetById(string) (*accessTokenService, *errors_utils.RestErr)
}

func NewAccessTokenService(repo Repository) accessTokenService {
	return accessTokenService{
		repository: repo,
	}
}

func (s *accessTokenService) GetById(ID string) (*accessTokenService, *errors_utils.RestErr) {

}
