package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConstantExpirationTime(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}
func TestGetNewAccessToken(t *testing.T) {
	at := &AccessToken{}

	assert.Equal(t, int64(0), at.Expires)

	assert.Equal(t, "", at.AccessToken)

	assert.NotEqual(t, 0, at.UserID)
}

func TestIsExpired(t *testing.T) {
	at := &AccessToken{}

	assert.Equal(t, false, at.IsExpired())

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.Equal(t, true, !at.IsExpired())
}
