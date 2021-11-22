package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mahdidl/simple_bank/db/sqlc"
)

//Server serves HTTP request fpr pur banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

//NewServer creates new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)


	server.router = router
	return server
}

//Start run HTTP server for specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
