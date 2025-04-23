package handler

import (
	"net/http"

	"github.com/first-restapi-golang/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) add(c echo.Context) error {
	var category model.Category

	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	return h.service.Add(&category)
}
func (h *Handler) getCategory(c echo.Context) error   { return nil }
func (h *Handler) getCategories(c echo.Context) error { return nil }
func (h *Handler) getTree(c echo.Context) error       { return nil }
func (h *Handler) changeParent(c echo.Context) error  { return nil }
func (h *Handler) exchangePlace(c echo.Context) error { return nil }
func (h *Handler) update(c echo.Context) error        { return nil }
func (h *Handler) delete(c echo.Context) error        { return nil }
