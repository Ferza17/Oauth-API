package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := RestUserRepositoryStruct{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid client response when trying to login user", err.Message)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": 404, "error": "not_found"}`,
	})
	repository := RestUserRepositoryStruct{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid client response when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": 404, "error": "not_found"}`,
	})
	repository := RestUserRepositoryStruct{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"id": "46","firstname": "Johny","lastname": "Silverhand","email": "johny@email.com"}`,
	})
	repository := RestUserRepositoryStruct{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when trying to unmarshal user response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "email@gmail.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"id": 46,"firstname": "Johny","lastname": "Silverhand","email": "johny@email.com"}`,
	})
	repository := RestUserRepositoryStruct{}
	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 46, user.ID)
	assert.EqualValues(t, "Johny",user.Firstname)
	assert.EqualValues(t, "Silverhand",user.Lastname)
	assert.EqualValues(t, "email@gmail.com",user.Email)
}
