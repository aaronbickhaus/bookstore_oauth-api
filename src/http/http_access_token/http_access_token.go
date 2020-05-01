package http_access_token

import (
	"github.com/aaronbickhaus/bookstore_oauth-api/src/domain/access_token"
	"github.com/aaronbickhaus/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}


func NewHandler(service access_token.Service) AccessTokenHandler {
   return &accessTokenHandler {
   	     service : service,
   }
}

func (h *accessTokenHandler) GetById (c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		e := errors.NewBadRequestError(err.Error())
		c.JSON(e.Status, e.Message)
		return
	}
	err := h.service.Create(at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, "created")
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		e := errors.NewBadRequestError(err.Error())
		c.JSON(e.Status, e.Message)
		return
	}
	err := h.service.UpdateExpriationTime(at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, "updated")
}