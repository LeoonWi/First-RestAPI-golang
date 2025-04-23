package handler

import (
	"github.com/first-restapi-golang/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func New(e *echo.Echo, s *service.Service) (h *Handler) {
	h = &Handler{s}
	category := e.Group("/category")
	{
		category.POST("/add", h.add)
		category.GET("/get", h.getCategories)
		category.GET("/get:id", h.getCategory)
		category.GET("/get-tree", h.getTree)
		category.PATCH("/change-parent", h.changeParent)
		category.PATCH("/exchange-place", h.exchangePlace)
		category.PUT("/update:id", h.update)
		category.DELETE("/delete:id", h.delete)

	}
	return
}
