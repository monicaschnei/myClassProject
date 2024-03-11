package api

import (
	"errors"
	"fmt"
	db "myclass/db/sqlc"
	"myclass/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createProfessionalInformationRequest struct {
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

type professionalUserId struct {
	UserName string `uri:"username" binding:"required"`
}

func (server *Server) createProfessionalInformation(ctx *gin.Context) {
	var req createProfessionalInformationRequest
	var reqUser professionalUserId

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, reqUser.UserName)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProfessionalInformationParams{
		ProfessionalUserID: professionalUser.ID,
		ExperiencePeriod:   req.ExperiencePeriod,
		OcupationArea:      req.OcupationArea,
		University:         req.University,
		GraduationDiploma:  req.GraduationDiploma,
		Validate:           false,
		GraduationCountry:  req.GraduationCountry,
		GraduationCity:     req.GraduationCity,
		GraduationState:    req.GraduationState,
		UpdatedAt:          time.Now(),
	}

	professionalInformation, err := server.store.CreateProfessionalInformation(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professionalInformation)
}

type getProfissionalInformationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProfessionalInformation(ctx *gin.Context) {
	var req getProfissionalInformationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalInformation, err := server.store.GetProfessionalInformation(ctx, req.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "Professional Information not found")
		return
	}
	ctx.JSON(http.StatusOK, professionalInformation)
}

type listProfessionalInformationsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAllProfessionalInformationsByUser(ctx *gin.Context) {
	var req listProfessionalInformationsRequest
	var reqUser professionalUserId

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, reqUser.UserName)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	professionalInformationsByUser, err := server.store.ListProfessionalInformationByUser(ctx, professionalUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, professionalInformationsByUser)
}

type updateProfessionalInformation struct {
	ExperiencePeriod  string `json:"experience_period"`
	OcupationArea     string `json:"ocupation_area"`
	University        string `json:"university"`
	GraduationDiploma string `json:"graduation_diploma"`
	Validate          bool   `json:"validate"`
	GraduationCountry string `json:"graduation_country"`
	GraduationCity    string `json:"graduation_city"`
	GraduationState   string `json:"graduation_state"`
}

func (server *Server) updateProfessionalInformation(ctx *gin.Context) {
	var req updateProfessionalInformation
	var reqUser professionalUserId
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, reqUser.UserName)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	arg := db.UpdateProfessionalInformationParams{
		ProfessionalUserID: professionalUser.ID,
		ExperiencePeriod:   req.ExperiencePeriod,
		OcupationArea:      req.OcupationArea,
		University:         req.University,
		GraduationDiploma:  req.GraduationDiploma,
		Validate:           false,
		GraduationCountry:  req.GraduationCountry,
		GraduationCity:     req.GraduationCity,
		GraduationState:    req.GraduationState,
	}
	professionalUserUpdated, err := server.store.UpdateProfessionalInformation(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not update this user")
		return
	}

	authPayload := ctx.MustGet(authorizatiionPayloadKey).(*token.Payload)
	if professionalUser.Username != authPayload.Username {
		err := errors.New("Account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, professionalUserUpdated)
}
