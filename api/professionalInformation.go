package api

import (
	db "myclass/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createProfessionalInformationRequest struct {
	ID                int64     `json:"id"`
	ExperiencePeriod  string    `json:"experience_period"`
	OcupationArea     string    `json:"ocupation_area"`
	University        string    `json:"university"`
	GraduationDiploma string    `json:"graduation_diploma"`
	Validate          bool      `json:"validate"`
	GraduationCountry string    `json:"graduation_country"`
	GraduationCity    string    `json:"graduation_city"`
	GraduationState   string    `json:"graduation_state"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (server *Server) createProfessionalInformation(ctx *gin.Context) {
	var req createProfessionalInformationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProfessionalInformationParams{
		ExperiencePeriod:  req.ExperiencePeriod,
		OcupationArea:     req.OcupationArea,
		University:        req.University,
		GraduationDiploma: req.GraduationDiploma,
		Validate:          req.Validate,
		GraduationCountry: req.GraduationCountry,
		GraduationCity:    req.GraduationCity,
		GraduationState:   req.GraduationState,
		UpdatedAt:         time.Now(),
	}

	professionalInformation, err := server.store.CreateProfessionalInformation(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, professionalInformation)
}
