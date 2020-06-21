package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := &AccessToken{}
	if at.IsExpired() {
		t.Error("new access token should now be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}

	if at.UserID != 0 {
		t.Error("new access token should not have an associated user id")
	}
}

func TestIsExpired(t *testing.T) {
	at := &AccessToken{}

	if at.IsExpired() {
		t.Error("new access token should not be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token created should not be expired after adding 3 hours")
	}
}
