package api

import (
	db "myclass/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/professionalUser", server.createProfessionalUser)
	router.PUT("/professionalUser/:id", server.updateProfessionalUser)
	router.GET("/professionalUser/:id", server.getProfessionalUserById)
	router.GET("/professionalUsers", server.listAllProfessionalUsers)
	router.DELETE("/professionalUser/:id", server.deleteProfessionalUser)
	router.POST("/professionalInformation/:id", server.createProfessionalInformation)
	router.GET("/professionalInformation/:id", server.getProfessionalInformation)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
