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
	//professionalUsers
	router.POST("/professionalUser", server.createProfessionalUser)
	router.PUT("/professionalUser/:id", server.updateProfessionalUser)
	router.GET("/professionalUser/:id", server.getProfessionalUserById)
	router.GET("/professionalUsers", server.listAllProfessionalUsers)
	router.DELETE("/professionalUser/:id", server.deleteProfessionalUser)
	//professionalInformations
	router.POST("/professionalInformation/:id", server.createProfessionalInformation)
	router.GET("/professionalInformation/:id", server.getProfessionalInformation)
	router.GET("/listProfessionalInformations/:id", server.listAllProfessionalInformationsByUser)
	router.PUT("/professionalInformations/:id", server.updateProfessionalInformation)
	//studentUser
	router.POST("/studentUser/:id", server.createStudentUser)
	router.GET("/studentUser/:id", server.getStudentUserById)
	router.GET("/listStudentUsers", server.listAllStudentUsers)
	router.PUT("/studentUser/:id", server.updateStudentlUser)
	//responsibleStudent
	router.POST("/responsibleStudent", server.createResponsibleStudent)
	router.GET("/responsibleStudent/:id", server.getResponsibleStudentUserById)
	router.GET("/listResponsibleStudent", server.listAllResponsibleStudents)
	router.PUT("/responsibleStudent/:id", server.updateResponsibleStudentl)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
