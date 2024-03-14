package service

import (
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"myclass/internal/core/model/request"
	"myclass/internal/core/model/response"
)

type ProfessionalUserService interface {
	CreateProfessionalUser(ctx *gin.Context, professionalUserRequest *request.CreateProfessionalUserRequest) (*response.ProfessionalUserResponse, error)
	GetProfessionalUser(ctx *gin.Context, username string) (db.ProfessionalUser, error)
	LoginProfessionalUser(ctx *gin.Context, request *request.LoginProfessionalUserRequest) (response.LoginProfessionalUserResponse, error)
}
