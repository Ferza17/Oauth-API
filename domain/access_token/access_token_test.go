package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T)  {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T)  {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t,"",at.AccessToken, "new Access Token should not have defined values")
	assert.True(t, at.UserId == 0,"New Access Token should not have associated user id")
}

func TestAccessTokenIsExpired( t *testing.T)  {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")
	if !at.IsExpired() {
		t.Error("Empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token created three hours from now should not be expired")
}