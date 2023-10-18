package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
	"net/http"
)

// CreateRating
// @Summary Create rating
// @Security ApiKeyAuth
// @Tags rating
// @Description create rating
// @ID create-rating
// @Accept  json
// @Produce  json
// @Param input body todo.RatingInput true "rating info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book/rating/ [post]
func (h *Handler) CreateRating(c *gin.Context) {
	var input todo.RatingInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Rating.CreateRating(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type GetRatingResponse struct {
	Rating []todo.GetRatingBook `json:"rating"`
}

// GetRatingList
// @Summary Get rating list
// @Security ApiKeyAuth
// @Tags rating
// @Description get rating list
// @ID get-rating-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} GetRatingResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book/rating/all [get]
func (h *Handler) GetRatingList(c *gin.Context) {
	ratingList, err := h.services.Rating.GetRatingList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetRatingResponse{
		Rating: ratingList,
	})
}
