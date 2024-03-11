package api

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "myclass/db/sqlc"
	"myclass/token"
	"myclass/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.AccesSymetryTokenKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gender", ValidGender)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwd", ValidPassword)
	}

	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//professionalUsers
	router.POST("/professionalUser", server.createProfessionalUser)
	router.POST("/professionalUser/login", server.loginProfessionalUser)
	authRoutes.PUT("/professionalUser/:username", server.updateProfessionalUser) //precisa estar loggado
	router.GET("/professionalUser/:username", server.getProfessionalUserById)
	router.GET("/professionalUsers", server.listAllProfessionalUsers)
	authRoutes.DELETE("/professionalUser/:username", server.deleteProfessionalUser) //precisa estar loggado

	//professionalInformations
	router.POST("/professionalInformation/:username", server.createProfessionalInformation)
	router.GET("/professionalInformation/:id", server.getProfessionalInformation)
	router.GET("/listProfessionalInformations/:id", server.listAllProfessionalInformationsByUser)
	authRoutes.PUT("/professionalInformations/:username", server.updateProfessionalInformation) //precisa estar loggado

	//studentUser
	router.POST("/studentUser/:username", server.createStudentUser)
	router.POST("/studentUser/login", server.loginStudentUser)
	router.GET("/studentUser/:username", server.getStudentUserById)
	router.GET("/listStudentUsers", server.listAllStudentUsers)
	authRoutes.PUT("/studentUser/:username/:responsibleStudent", server.updateStudentlUser) //precisa estar loggado

	//responsibleStudent
	router.POST("/responsibleStudent", server.createResponsibleStudent)
	router.POST("/responsibleStudent/login", server.loginResponsibleStudent)
	router.GET("/responsibleStudent/:username", server.getResponsibleStudentUserById)
	router.GET("/listResponsibleStudent", server.listAllResponsibleStudents)
	authRoutes.PUT("/responsibleStudent/:username", server.updateResponsibleStudent) //precisa estar loggado

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
