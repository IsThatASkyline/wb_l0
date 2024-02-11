package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOrderByUid(c *gin.Context) {
	uid := c.Params.ByName("id")

	order, err := h.services.Order.GetOrderByUid(uid)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
