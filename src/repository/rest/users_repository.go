package rest

import (
	"encoding/json"
	"time"

	"github.com/flucas97/bookstore/users-api/model/users"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	Login(string, string) (*users.User, *errors_utils.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) Login(email string, password string) (*users.User, *errors_utils.RestErr) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := restClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors_utils.NewInternalServerError("invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors_utils.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors_utils.NewInternalServerError("invalid error when trying to login")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors_utils.NewNotFoundError("error when trying to unmarshal users response")
	}

	return &user, nil
}
