package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAcessTokenConstants(t *testing.T) {
  assert.EqualValues(t, 24, expirationTime, "expiration time for token should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
    assert.False(t, at.IsExpired(), "new token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "new token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new token should not have defined user id")

}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")


	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token should not be expired")

}
