package api

import (
	db "github.com/nadeeracode/GoLoginSvc/db/sqlc"
	"github.com/nadeeracode/GoLoginSvc/util"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking server
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router

}

// Start  runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
