package rest

import (
	"encoding/json"
	"github.com/Ferza17/Oauth-API/src/domain/user"
	"github.com/Ferza17/Oauth-API/src/utils/errors"
	"github.com/Ferza17/Oauth-API/domain/user"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,

	}
)

type RestUserRepositoryInterface interface {
	LoginUser(string, string) (*user.User, *errors.RestError)
}

type RestUserRepositoryStruct struct {
}

func (r *RestUserRepositoryStruct) LoginUser(email string, password string) (*user.User, *errors.RestError) {
	request := users.LoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil{
		return nil, errors.NewInternalServerError("Invalid client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestError
		 err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to logging user")
		}
		return nil, &restErr
	}
	
	var result user.User
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal user response")
	}

	return &result, nil
}

func NewRepository() *RestUserRepositoryStruct {
	return &RestUserRepositoryStruct{}
}
