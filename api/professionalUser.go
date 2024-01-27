package api

import (
	db "myclass/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type createProfessionalUserRequest struct {
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Gender         string    `json:"gender"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Cpf            int32     `json:"cpf"`
	ClassHourPrice string    `json:"class_hour_price"`
}

func (server *Server) createProfessionalUser(ctx *gin.Context) {
	var req createProfessionalUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProfessionalUserParams{
		Name:                      req.Name,
		Username:                  req.Username,
		Password:                  req.Password,
		Gender:                    req.Gender,
		Email:                     req.Email,
		DateOfBirth:               req.DateOfBirth,
		Cpf:                       req.Cpf,
		ImageID:                   0,
		PhoneID:                   0,
		ProfessionalInformationID: 0,
		UpdatedAt:                 time.Now(),
		SubjectMatterID:           0,
		SubjectMatterClassID:      0,
		ClassHourPrice:            req.ClassHourPrice,
		CalendarID:                0,
	}

	professionalUser, err := server.store.CreateProfessionalUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, professionalUser)
}
