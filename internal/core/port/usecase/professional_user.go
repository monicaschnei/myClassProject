package usecase

import (
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"myclass/internal/core/domain/request"
	"myclass/internal/core/domain/response"
	"myclass/token"
)

type ProfessionalUserService interface {
	CreateProfessionalUser(ctx *gin.Context, professionalUserRequest *request.CreateProfessionalUserRequest) (*response.ProfessionalUserResponse, error)
	GetProfessionalUser(ctx *gin.Context, username string) (db.ProfessionalUser, error)
	LoginProfessionalUser(ctx *gin.Context, request *request.LoginProfessionalUserRequest) (response.LoginProfessionalUserResponse, error)
	ListProfessionalUser(ctx *gin.Context, request *request.ListProfessionalUsersRequest) ([]db.ProfessionalUser, error)
	UpdateProfessionalUser(ctx *gin.Context, authPayload *token.Payload, username string, request *request.UpdateProfessionalUserRequest) (response.ProfessionalUserResponse, error)
}
