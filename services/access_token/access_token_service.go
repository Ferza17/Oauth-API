package access_token_service

import (
	"github.com/Ferza17/Oauth-API/src/domain/access_token"
	"github.com/Ferza17/Oauth-API/src/repository/db"
	"github.com/Ferza17/Oauth-API/src/repository/rest"
	"github.com/Ferza17/Oauth-API/src/utils/errors"
	"strings"
)

type ServiceInterface interface {
	GetById(atId string) (*access_token.AccessToken, *errors.RestError)
	Create(request access_token.AccessTokenRequest) (*access_token.AccessToken,*errors.RestError)
	UpdateExpirationTime(at access_token.AccessToken) *errors.RestError
}

type serviceStruct struct {
	restUserRepo rest.RestUserRepositoryInterface
	dbRepo db.DbRepositoryInterface
}

func NewService(userRepo rest.RestUserRepositoryInterface, dbRepo db.DbRepositoryInterface) ServiceInterface {
	return &serviceStruct{
		restUserRepo: userRepo,
		dbRepo: dbRepo,
	}
}

func (s *serviceStruct) GetById (atId string)(*access_token.AccessToken, *errors.RestError) {
	accessTokenId := strings.TrimSpace(atId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid Access Token ID")
	}
	accessToken, err :=  s.dbRepo.GetById(atId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil

}

func (s *serviceStruct) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken,*errors.RestError)  {
	if err:= request.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}


	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	//Save the new Access token to cassandra db
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *serviceStruct) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError  {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}

