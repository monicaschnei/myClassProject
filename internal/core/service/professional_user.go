package service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-cpf"
	"myclass/api"
	db "myclass/db/sqlc"
	"myclass/internal/core/model/request"
	"myclass/internal/core/model/response"
	"myclass/internal/core/port/service"
	"myclass/util"
	"time"
)

type professionalUserService struct {
	userRepo db.Store
	server   api.HttpServer
}

func NewProfessionalUserService(userRepo db.Store, server api.HttpServer) service.ProfessionalUserService {
	return &professionalUserService{
		userRepo: userRepo,
		server:   server,
	}
}

func (pus professionalUserService) CreateProfessionalUser(ctx *gin.Context, request *request.CreateProfessionalUserRequest) (*response.ProfessionalUserResponse, error) {
	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return &response.ProfessionalUserResponse{}, err
	}

	dateOfBirth, err := util.TransformDateOfBirth("2006-01-02", request.DateOfBirth)
	if err != nil {
		return &response.ProfessionalUserResponse{}, err
	}

	if validCpf, err := cpf.Valid(request.Cpf); !validCpf {
		if err != nil {
			return &response.ProfessionalUserResponse{}, fmt.Errorf("Internal Server Error")
		}
		return &response.ProfessionalUserResponse{}, fmt.Errorf("Invalid cpf")
	}

	userDb := db.CreateProfessionalUserParams{
		Name:           request.Name,
		Username:       request.Username,
		HashedPassword: hashedPassword,
		Gender:         request.Gender,
		Email:          request.Email,
		DateOfBirth:    dateOfBirth,
		Cpf:            request.Cpf,
		ImageID:        request.ImageID,
		UpdatedAt:      time.Now(),
		ClassHourPrice: request.ClassHourPrice,
	}

	professionalUser, err := pus.userRepo.CreateProfessionalUser(ctx, userDb)
	response := response.NewProfessionalUserResponse(professionalUser)
	return &response, nil
}

func (pus professionalUserService) GetProfessionalUser(ctx *gin.Context, username string) (db.ProfessionalUser, error) {
	professionalUser, err := pus.userRepo.GetProfessionalUser(ctx, username)
	if err != nil {
		return db.ProfessionalUser{}, fmt.Errorf("This user does not exists, please create it firstly")
	}
	return professionalUser, nil
}

func (pus professionalUserService) LoginProfessionalUser(ctx *gin.Context, request *request.LoginProfessionalUserRequest) (response.LoginProfessionalUserResponse, error) {
	professionalUser, err := pus.GetProfessionalUser(ctx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.LoginProfessionalUserResponse{}, fmt.Errorf("There is no professional user with this username %s", request.Username)
		}
		return response.LoginProfessionalUserResponse{}, fmt.Errorf("Internal Server Error")
	}

	err = util.CheckPassword(request.Password, professionalUser.HashedPassword)
	if err != nil {
		return response.LoginProfessionalUserResponse{}, fmt.Errorf("Unauthorized, wrong password")
	}

	tokenMaker, _ := pus.server.TokenMaker()
	accesToken, err := tokenMaker.CreateToken(
		professionalUser.Username,
		pus.server.AccessTokenDuration(),
	)

	response := pus.NewProfesionalUserLoginResponse(accesToken, professionalUser)
	return response, nil
}

func (pus professionalUserService) NewProfesionalUserLoginResponse(acessToken string, professionalUser db.ProfessionalUser) response.LoginProfessionalUserResponse {
	return response.LoginProfessionalUserResponse{
		AccessToken: acessToken,
		User:        response.NewProfessionalUserResponse(professionalUser),
	}
}
