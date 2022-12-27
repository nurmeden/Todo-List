package handler

import (
	"html/template"
	"net/http"
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler { // конструктор
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine { // endpoint
	router := gin.New() // initialize router

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.GET("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}

	return router
}

func Create_Page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("webpage.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
