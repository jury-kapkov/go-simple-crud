package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {

}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) getList(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
