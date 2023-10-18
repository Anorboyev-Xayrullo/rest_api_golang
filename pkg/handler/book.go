package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
	"net/http"
	"strconv"
)

// CreateBook
// @Summary Create Book
// @Security ApiKeyAuth
// @Tags book
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body todo.CreateBook true "book info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var input todo.CreateBook
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Book.CreateBook(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type GetBookListResponse struct {
	Book []todo.GetBookList `json:"book"`
}

// GetBookList
// @Summary Get Book list
// @Security ApiKeyAuth
// @Tags book
// @Description get book list
// @ID get-book-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} GetBookListResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book/list [get]
func (h *Handler) GetBookList(c *gin.Context) {
	bookList, err := h.services.Book.GetBookList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetBookListResponse{
		Book: bookList,
	})
}

// GetBookById
// @Summary Get Book By Id
// @Security ApiKeyAuth
// @Tags book
// @Description get book by id
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} todo.SuccessData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book/id [get]
func (h *Handler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	list, err := h.services.Book.GetBookById(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo.SuccessData{
		Data: list,
	})
}

// UpdateBook
// @Summary update Book
// @Security ApiKeyAuth
// @Tags book
// @Description Update Book
// @ID update book
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Param input body todo.CreateBook true "book info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Book.UpdateBookInput(uint(bookId), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// DeleteBook
// @Summary Delete Book
// @Security ApiKeyAuth
// @Tags book
// @Description delete book
// @ID delete book
// @Accept  json
// @Produce  json
// @Param id query int true "id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/book [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Book.DeleteBook(uint(bookId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
