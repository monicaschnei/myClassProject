package usecase

import (
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"myclass/internal/core/domain/request"
	"myclass/internal/core/domain/response"
)

type ProfessionalAvailabilityService interface {
	AddAvailability(ctx *gin.Context, request request.AddAvailabilityRequest) (response.AddAvailabilityResponse, error)
	ListAvailabilityFromProfessionalUser(ctx *gin.Context, userName string, request request.ListAvailabilityRequest) ([]db.Availability, error)
}
