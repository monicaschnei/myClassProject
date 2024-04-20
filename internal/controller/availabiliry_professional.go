package controller

import (
	"github.com/gin-gonic/gin"
	"myclass/internal/core/domain/request"
	"myclass/internal/core/port/usecase"
	"myclass/internal/infrastructure/router"
	"net/http"
)

type AvailabilityProfessionalUserController struct {
	gin                      *gin.Engine
	availabilityProfessional usecase.ProfessionalAvailabilityService
}

func NewAvailabilityProfessionalUserController(gin *gin.Engine, availabilityProfessional usecase.ProfessionalAvailabilityService) AvailabilityProfessionalUserController {
	return AvailabilityProfessionalUserController{
		gin:                      gin,
		availabilityProfessional: availabilityProfessional,
	}
}
func (a AvailabilityProfessionalUserController) InitRouter() {
	api := a.gin.Group("/")

	router.Post(api, "/addAvailability", a.addAvailability)
	router.Get(api, "/addAvailability/:username", a.listAvailability)
}

func (a AvailabilityProfessionalUserController) addAvailability(ctx *gin.Context) {
	var req request.AddAvailabilityRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availabilityAdded, err := a.availabilityProfessional.AddAvailability(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, availabilityAdded)
}

func (a AvailabilityProfessionalUserController) listAvailability(ctx *gin.Context) {
	var req request.ListAvailabilityRequest
	var reqUser request.GetProfissionalUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listAvailability, err := a.availabilityProfessional.ListAvailabilityFromProfessionalUser(ctx, reqUser.UserName, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listAvailability)
}
