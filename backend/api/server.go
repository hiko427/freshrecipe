package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hiko427/myrecipe/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func Newserver(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.LoadHTMLGlob("../frontend/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	router.POST("/submit", server.createRecipe)
	router.GET("/show/:recipeID", server.showRecipe)
	router.GET("/history", func(c *gin.Context) {
		c.HTML(http.StatusOK, "history.tmpl", nil)
	})
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
