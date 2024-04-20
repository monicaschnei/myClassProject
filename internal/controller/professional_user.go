package controller

import (
	"github.com/gin-gonic/gin"
	"myclass/api"
	"myclass/internal/core/domain/request"
	"myclass/internal/core/port/usecase"
	"myclass/internal/infrastructure/router"
	"myclass/token"
	"net/http"
)

type ProfessionalUserController struct {
	gin         *gin.Engine
	userService usecase.ProfessionalUserService
	tokenMaker  token.Maker
}

func NewProfessionalUserController(gin *gin.Engine, userService usecase.ProfessionalUserService) ProfessionalUserController {
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
	router.Post(api, "/professionalUsers", u.listAllProfessionalUsers)
	router.Put(authApi, "/professionalUser/:username", u.updateProfessionalUser)
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

func (u ProfessionalUserController) listAllProfessionalUsers(ctx *gin.Context) {
	var req *request.ListProfessionalUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	professionalUsers, err := u.userService.ListProfessionalUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, professionalUsers)
}

func (u ProfessionalUserController) updateProfessionalUser(ctx *gin.Context) {
	var req request.UpdateProfessionalUserRequest
	var reqUser request.GetProfissionalUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var authPayload = ctx.MustGet("authorization_payload").(*token.Payload)
	professionalUserUpdated, err := u.userService.UpdateProfessionalUser(ctx, authPayload, reqUser.UserName, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professionalUserUpdated)
}
