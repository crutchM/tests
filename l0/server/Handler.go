package server

import (
	"github.com/gin-gonic/gin"
	"tests/l0/repo"
)

type Handler struct {
	repo *repo.Repository
}

func NewHandler(repo *repo.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:id", h.getOrder)

	return router
}

func (h *Handler) getOrder(c *gin.Context) {
	id := c.Param("id")
	item := h.repo.Get(id)
	c.JSON(200, map[string]interface{}{"order": item})
}
