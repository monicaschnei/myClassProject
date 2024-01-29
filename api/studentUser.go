package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"net/http"
	"time"
)

type createStudentUserRequest struct {
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	Name                 string    `json:"name"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	Gender               string    `json:"gender"`
	ResponsibleStudentID int32     `json:"responsible_student_id"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (server *Server) createStudentUser(ctx *gin.Context) {
	var req createStudentUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStudentUserParams{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		UpdatedAt:   req.UpdatedAt,
	}

	studentUser, err := server.store.CreateStudentUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, studentUser)
}

type getStudentlUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStudentUserById(ctx *gin.Context) {
	var req getStudentlUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	studentUser, err := server.store.GetStudentUser(ctx, req.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	ctx.JSON(http.StatusOK, studentUser)
}

type listStudentUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAllStudentUsers(ctx *gin.Context) {
	var req listStudentUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStudentUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	studentUsers, err := server.store.ListStudentUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, studentUsers)
}

type updateStudentlUser struct {
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	Name                 string    `json:"name"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	Gender               string    `json:"gender"`
	ResponsibleStudentID int32     `json:"responsible_student_id"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (server *Server) updateStudentlUser(ctx *gin.Context) {
	var req updateStudentlUser
	var reqID getStudentlUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	studentUser, err := server.store.GetStudentUser(ctx, reqID.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	arg := db.UpdateStudentUserParams{
		ID:       studentUser.ID,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
	studentUserUpdated, err := server.store.UpdateStudentUser(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not update this user")
		return
	}
	ctx.JSON(http.StatusOK, studentUserUpdated)
}
