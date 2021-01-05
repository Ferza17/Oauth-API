package db

import (
	"github.com/Ferza17/Oauth-API/src/clients/cassandra"
	"github.com/Ferza17/Oauth-API/src/domain/access_token"
	"github.com/Ferza17/Oauth-API/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token=?"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, client_id, expires, user_id) VALUES (?,?,?,?)"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?"
	)

type DbRepositoryInterface interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestError)
	Create(at access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(at access_token.AccessToken) *errors.RestError
}

type DbRepositoryStruct struct {}

func NewRepository() *DbRepositoryStruct {
	return &DbRepositoryStruct{}
}

func (r *DbRepositoryStruct) GetById(id string) (*access_token.AccessToken, *errors.RestError)  {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.ClientId,
		&result.Expires,
		&result.UserId); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *DbRepositoryStruct) Create(at access_token.AccessToken) *errors.RestError  {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.ClientId,
		at.Expires,
		at.UserId).Exec(); err != nil{
		return errors.NewInternalServerError(err.Error())
	}

	return nil

}

func (r *DbRepositoryStruct) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError  {

	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}