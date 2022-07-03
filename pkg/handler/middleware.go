package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo/pkg/messages"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentifier(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, messages.EmptyAuthHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, messages.InvalidAuthHeader)
		return
	}

	userId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, messages.InvalidAuthHeader)
		return 0, errors.New(messages.InvalidAuthHeader)
	}

	intId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, messages.InvalidTypeOfUserId)
		return 0, errors.New(messages.InvalidTypeOfUserId)
	}

	return intId, nil
}
