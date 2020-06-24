package access_token

import (
	"strings"
	"time"

	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

const (
	expirationTime = 24
)

// model
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// model functions
func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// validations
func (at *AccessToken) IsExpired() bool {
	return false
}

func (at *AccessToken) Validate() *errors_utils.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors_utils.NewBadRequestError("invalid access token")
	}

	if at.UserID < 0 {
		return errors_utils.NewBadRequestError("invalid user ID")
	}

	if at.ClientID < 0 {
		return errors_utils.NewBadRequestError("invalid client ID")
	}

	if at.Expires < 0 {
		return errors_utils.NewBadRequestError("invalid expiration time")
	}
	return nil
}
