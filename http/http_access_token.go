package http

import (
	"github.com/Ferza17/Oauth-API/src/domain/access_token"
	access_token_service "github.com/Ferza17/Oauth-API/src/services/access_token"
	"github.com/Ferza17/Oauth-API/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandlerInterface interface {
	GetById(*gin.Context)
	Create(ctx *gin.Context)
}

type AccessTokenHandlerStruct struct {
	service access_token_service.ServiceInterface
}
func NewAccessTokenHandler(service access_token_service.ServiceInterface) *AccessTokenHandlerStruct {
	return &AccessTokenHandlerStruct{
		service: service,
	}
}

func (h *AccessTokenHandlerStruct) GetById(c *gin.Context) {
	accessToken, err:= h.service.GetById(c.Param("access_token_id"))
	if  err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *AccessTokenHandlerStruct) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}

// TODO : Update expiration Time
//func (h *AccessTokenHandlerStruct) UpdateExpirationTime (c *gin.Context){
//
//}
