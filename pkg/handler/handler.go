package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/zhashkevych/todo-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	apiv1 := router.Group("api/v1")
	{
		book := apiv1.Group("/book")
		{
			book.POST("/", h.CreateBook)
			book.GET("/list", h.GetBookList)
			book.GET("/id", h.GetBookById)
			book.PUT("/", h.UpdateBook)
			book.DELETE("/", h.DeleteBook)
			rating := book.Group("/rating")
			{
				rating.POST("/", h.CreateRating)
				rating.GET("/all", h.GetRatingList)
			}
		}
		gener := apiv1.Group("/gener")
		{
			gener.POST("/", h.CreateGener)
			gener.GET("/", h.GetGenerList)
			gener.GET("/id", h.GetGenerById)
			gener.PUT("/", h.UpdateGener)
			gener.DELETE("/", h.DeleteGener)
		}
		library := apiv1.Group("/library")
		{
			library.GET("/", h.GetLibraryList)
		}

	}

	return router
}
