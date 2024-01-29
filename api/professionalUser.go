package api

import (
	"fmt"
	db "myclass/db/sqlc"
	"myclass/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ProfessionalUserController struct {
	professionalUserService *services.ProfessionalUserService
}

type createProfessionalUserRequest struct {
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Gender         string    `json:"gender"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Cpf            int32     `json:"cpf"`
	ClassHourPrice string    `json:"class_hour_price"`
	ImageID        int64     `json:"image_id"`
}

func (server *Server) createProfessionalUser(ctx *gin.Context) {
	var req createProfessionalUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProfessionalUserParams{
		Name:           req.Name,
		Username:       req.Username,
		Password:       req.Password,
		Gender:         req.Gender,
		Email:          req.Email,
		DateOfBirth:    req.DateOfBirth,
		Cpf:            req.Cpf,
		ImageID:        req.ImageID,
		UpdatedAt:      time.Now(),
		ClassHourPrice: req.ClassHourPrice,
	}

	professionalUser, err := server.store.CreateProfessionalUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, professionalUser)
}

type getProfissionalUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProfessionalUserById(ctx *gin.Context) {
	var req getProfissionalUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, req.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	ctx.JSON(http.StatusOK, professionalUser)
}

type listProfessionalUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAllProfessionalUsers(ctx *gin.Context) {
	var req listProfessionalUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProfessionalUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	professionalUsers, err := server.store.ListProfessionalUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, professionalUsers)
}

type updateProfessionalUser struct {
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	ClassHourPrice string    `json:"class_hour_price"`
}

func (server *Server) updateProfessionalUser(ctx *gin.Context) {
	var req updateProfessionalUser
	var reqID getProfissionalUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, reqID.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	arg := db.UpdateProfessionalUserParams{
		ID:             professionalUser.ID,
		Name:           req.Name,
		Username:       req.Username,
		Password:       req.Password,
		Email:          req.Email,
		DateOfBirth:    req.DateOfBirth,
		ClassHourPrice: req.ClassHourPrice,
	}
	professionalUserUpdated, err := server.store.UpdateProfessionalUser(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not update this user")
		return
	}
	ctx.JSON(http.StatusOK, professionalUserUpdated)
}

func (server *Server) deleteProfessionalUser(ctx *gin.Context) {
	var reqID getProfissionalUserRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	professionalUser, err := server.store.GetProfessionalUser(ctx, reqID.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exist")
		return
	}

	_, err = server.store.DeleteProfessionalUser(ctx, professionalUser.ID)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not delete this user")
		return
	}
}
