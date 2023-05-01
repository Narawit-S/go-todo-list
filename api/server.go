package api

import (
	db "github.com/Narawit-S/go-todo-list/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	///api
	router.POST("/register", server.createUser)

	server.router = router
	return server
}

func (server *Server) Start(port string) error {
	return server.router.Run(port)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
