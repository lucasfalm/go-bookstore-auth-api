package access_token

import (
	"testing"
)

func TestGetNewAccessToken(t *testing.T) {
	at := &AccessToken{}
	if at.IsExpired() {
		t.Error("new access token should now be expired")
	}
}
