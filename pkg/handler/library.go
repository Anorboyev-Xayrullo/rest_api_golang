package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
	"net/http"
)

type GetLibraryResponse struct {
	Library []todo.Library `json:"library"`
}

// GetLibraryList
// @Summary Get Library
// @Security ApiKeyAuth
// @Tags library
// @Description get library
// @ID get-library
// @Accept  json
// @Produce  json
// @Success 200 {object} GetLibraryResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/library/ [get]
func (h *Handler) GetLibraryList(c *gin.Context) {
	librarylist, err := h.services.Library.GetLibraryList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetLibraryResponse{
		Library: librarylist,
	})
}
