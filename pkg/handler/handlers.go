package handler

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) initRoutes() *gin.Engine { // endpoint
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
			lists.POST("/")
		}
	}
}

func Create_Page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("webpage.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
