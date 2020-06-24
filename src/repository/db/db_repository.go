package db

import (
	"fmt"

	"github.com/flucas97/bookstore/auth-api/clients/cassandra"
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
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
	accessToken := access_token.AccessToken{}
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors_utils.NewInternalServerError(fmt.Sprintf("error: %v", err.Error()))
	}
	defer session.Close()

	if err := session.Query(queryGetAccessToken, id).Scan(
		&accessToken.AccessToken,
		&accessToken.UserID,
		&accessToken.ClientID,
		&accessToken.Expires,
	); err != nil {
		return nil, errors_utils.NewInternalServerError(fmt.Sprintf("error: %v", err.Error()))
	}

	return &accessToken, nil
}
