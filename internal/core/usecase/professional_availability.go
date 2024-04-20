package usecase

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"myclass/api"
	db "myclass/db/sqlc"
	"myclass/internal/core/domain/request"
	"myclass/internal/core/domain/response"
	"myclass/internal/core/port/usecase"
)

type professionalAvailabilityService struct {
	userRepo                db.Store
	server                  api.HttpServer
	professionalUserService usecase.ProfessionalUserService
}

func NewProfessionalAvailabilityService(userRepo db.Store, server api.HttpServer, professionalUserService usecase.ProfessionalUserService) usecase.ProfessionalAvailabilityService {
	return &professionalAvailabilityService{
		userRepo:                userRepo,
		server:                  server,
		professionalUserService: professionalUserService,
	}
}
func (pas professionalAvailabilityService) AddAvailability(ctx *gin.Context, request request.AddAvailabilityRequest) (response.AddAvailabilityResponse, error) {
	professionalUser, err := pas.professionalUserService.GetProfessionalUser(ctx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.AddAvailabilityResponse{}, fmt.Errorf("There is no professional user with this username %s", request.Username)
		}
		return response.AddAvailabilityResponse{}, fmt.Errorf("Internal Server Error")
	}

	addVailability := db.AddAvailabilityParams{
		Date:        request.Date,
		UserID:      professionalUser.ID,
		Username:    professionalUser.Username,
		Start:       request.Start,
		EndTime:     request.EndTime,
		IsAvailable: request.IsAvailable,
	}
	availabilityAdded, err := pas.userRepo.AddAvailability(ctx, addVailability)
	response := response.NewAvaialibilityResponse(availabilityAdded)
	return response, nil
}

func (pas professionalAvailabilityService) ListAvailabilityFromProfessionalUser(ctx *gin.Context, userName string, request request.ListAvailabilityRequest) ([]db.Availability, error) {
	_, err := pas.professionalUserService.GetProfessionalUser(ctx, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return []db.Availability{}, fmt.Errorf("There is no professional user with this username %s", userName)
		}
		return []db.Availability{}, fmt.Errorf("Internal Server Error")
	}

	listParams := db.ListAvailabilityParams{
		Limit:  request.PageSize,
		Offset: (request.PageID - 1) * request.PageSize,
	}

	listAvailability, err := pas.userRepo.ListAvailability(ctx, listParams)
	return listAvailability, err
}

func (pas professionalAvailabilityService) NewProfessionalAvailabilityResponse(professionalUser db.ProfessionalUser, date, start, endTime string, isAvaialable bool) response.AddAvailabilityResponse {
	return response.AddAvailabilityResponse{
		Date:        date,
		UserId:      professionalUser.ID,
		Username:    professionalUser.Username,
		Start:       start,
		EndTime:     endTime,
		IsAvailable: isAvaialable,
	}
}
