package api

import (
	"fmt"
	db "myclass/db/sqlc"
	"myclass/token"
	"myclass/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
	Start() error
	TokenMaker() (token.Maker, error)
	AccessTokenDuration() time.Duration
}

type httpServer struct {
	Port       uint
	server     *http.Server
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	//router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(router *gin.Engine, config util.Config) HttpServer {

	return httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
			Handler: router,
		},
		config: config,
	}

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("gender", ValidGender)
	//}
	//
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("passwd", ValidPassword)
	//}

	//router := gin.Default()
	//server.router = router
	//server.setUpRouter()
	//return server, nil
}

//func (server *Server) setUpRouter() {
//	router := gin.Default()
//
//	authRoutes := router.Group("/").Use(AuthMiddleware(server.tokenMaker))
//
//	//professionalUsers
//	router.POST("/professionalUser", server.createProfessionalUser)
//	router.POST("/professionalUser/login", server.loginProfessionalUser)
//	authRoutes.PUT("/professionalUser/:username", server.updateProfessionalUser) //precisa estar loggado
//	router.GET("/professionalUser/:username", server.getProfessionalUserById)
//	router.GET("/professionalUsers", server.listAllProfessionalUsers)
//	authRoutes.DELETE("/professionalUser/:username", server.deleteProfessionalUser) //precisa estar loggado
//
//	//professionalInformations
//	router.POST("/professionalInformation/:username", server.createProfessionalInformation)
//	router.GET("/professionalInformation/:id", server.getProfessionalInformation)
//	router.GET("/listProfessionalInformations/:id", server.listAllProfessionalInformationsByUser)
//	authRoutes.PUT("/professionalInformations/:username", server.updateProfessionalInformation) //precisa estar loggado
//
//	//studentUser
//	router.POST("/studentUser/:username", server.createStudentUser)
//	router.POST("/studentUser/login", server.loginStudentUser)
//	router.GET("/studentUser/:username", server.getStudentUserById)
//	router.GET("/listStudentUsers", server.listAllStudentUsers)
//	authRoutes.PUT("/studentUser/:username/:responsibleStudent", server.updateStudentlUser) //precisa estar loggado
//
//	//responsibleStudent
//	router.POST("/responsibleStudent", server.createResponsibleStudent)
//	router.POST("/responsibleStudent/login", server.loginResponsibleStudent)
//	router.GET("/responsibleStudent/:username", server.getResponsibleStudentUserById)
//	router.GET("/listResponsibleStudent", server.listAllResponsibleStudents)
//	authRoutes.PUT("/responsibleStudent/:username", server.updateResponsibleStudent) //precisa estar loggado
//
//	server.router = router
//}

func (httpServer httpServer) Start() error {
	return httpServer.server.ListenAndServe()
}

func (httpServer httpServer) TokenMaker() (token.Maker, error) {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, err
		//fmt.Errorf("Cannot create token maker: %w", err)
	}
	return tokenMaker, nil
}

func (httpServer httpServer) AccessTokenDuration() time.Duration {
	return httpServer.config.AccessTokenDuration
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
