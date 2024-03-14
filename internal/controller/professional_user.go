package controller

import (
	"github.com/gin-gonic/gin"
	"myclass/api"
	"myclass/internal/core/common/router"
	"myclass/internal/core/model/request"
	"myclass/internal/core/port/service"
	"myclass/token"
	"net/http"
)

type ProfessionalUserController struct {
	gin         *gin.Engine
	userService service.ProfessionalUserService
	tokenMaker  token.Maker
}

func NewProfessionalUserController(gin *gin.Engine, userService service.ProfessionalUserService) ProfessionalUserController {
	return ProfessionalUserController{
		gin:         gin,
		userService: userService,
	}
}
func (u ProfessionalUserController) InitRouter() {
	authApi := u.gin.Group("/")
	authApi.Use(api.AuthMiddleware(u.tokenMaker))
	api := u.gin.Group("/")

	router.Post(api, "/professionalUser", u.createProfessionalUser)
	router.Get(api, "/professionalUser/:username", u.getProfessionalUser)
	router.Post(api, "/professionalUser/login", u.loginProfessionalUser)
}

func (u ProfessionalUserController) createProfessionalUser(ctx *gin.Context) {
	var req *request.CreateProfessionalUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	professionalUser, err := u.userService.CreateProfessionalUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, professionalUser)
}

func (u ProfessionalUserController) getProfessionalUser(ctx *gin.Context) {
	var req *request.GetProfissionalUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	professionalUser, err := u.userService.GetProfessionalUser(ctx, req.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professionalUser)
}

func (u ProfessionalUserController) loginProfessionalUser(ctx *gin.Context) {
	var req *request.LoginProfessionalUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	professionalUserLogged, err := u.userService.LoginProfessionalUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professionalUserLogged)
}
