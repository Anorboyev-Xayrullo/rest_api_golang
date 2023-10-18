package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
	_ "github.com/zhashkevych/todo-app"
	"net/http"
	"strconv"
)

// CreateGener
// @Summary Create gener
// @Security ApiKeyAuth
// @Tags gener
// @Description create gener
// @ID create-gener
// @Accept  json
// @Produce  json
// @Param input body todo.Gener true "genre info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/gener/ [post]
func (h *Handler) CreateGener(c *gin.Context) {
	var input todo.Gener
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Gener.CreateGener(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type GetGenerResponse struct {
	Gener []todo.GetGenerBook `json:"generall"`
}

// GetGenerList
// @Summary Get Gener
// @Security ApiKeyAuth
// @Tags gener
// @Description get gener
// @ID get-gener
// @Accept  json
// @Produce  json
// @Success 200 {object} GetGenerResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/gener/ [get]
func (h *Handler) GetGenerList(c *gin.Context) {
	generlist, err := h.services.Gener.GetGenerList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetGenerResponse{
		Gener: generlist,
	})
}

// GetGenerById
// @Summary Get Gener By id
// @Security ApiKeyAuth
// @Tags gener
// @Description get gener by id
// @ID get-gener-by-id
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} todo.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/gener/id  [get]
func (h *Handler) GetGenerById(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.Gener.GetGenerById(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo.SuccessData{
		Data: list,
	})
}

// UpdateGener
// @Summary update gener
// @Security ApiKeyAuth
// @Tags gener
// @Description update gener
// @ID update gener
// @Accept  json
// @Produce  json
// @Param id query int true "ID"
// @Param input body todo.UpdateGenerInput true "book info"
// @Success 200 {object} todo.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/gener/ [put]
func (h *Handler) UpdateGener(c *gin.Context) {
	generId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateGenerInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Gener.UpdateGenerInput(uint(generId), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// DeleteGener
// @Summary delete gener
// @Security ApiKeyAuth
// @Tags gener
// @Description delete gener
// @ID delete gener
// @Accept  json
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {object} todo.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/gener/ [delete]
func (h *Handler) DeleteGener(c *gin.Context) {

	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Gener.DeleteGener(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo.SuccessData{
		Data: "ok",
	})
}
