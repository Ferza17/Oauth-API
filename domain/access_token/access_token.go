package access_token

import (
	"fmt"
	"github.com/Ferza17/Oauth-API/src/utils/errors"
	"github.com/Ferza17/Oauth-API/utils/crypt"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestError{
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("Invalid grant_type parameter")
	}
	// TODO: validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid token id")
	}

	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}

	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func (at *AccessToken) Generate(){
	at.AccessToken =crypt.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

// Web Frontend - Client-id: 123
// Android APP - Client-id: 234
