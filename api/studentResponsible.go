package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"net/http"
	"time"
)

type createStudentResponsbileRequest struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	Cpf         int32     `json:"Cpf"`
	Gender      string    `json:"gender"`
}

func (server *Server) createResponsibleStudent(ctx *gin.Context) {
	var req createStudentResponsbileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateResponsibleStudentParams{
		Name:        req.Name,
		Gender:      req.Gender,
		Email:       req.Email,
		DateOfBirth: req.DateOfBirth,
		Cpf:         req.Cpf,
	}

	responsibleStudent, err := server.store.CreateResponsibleStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, responsibleStudent)
}

type getResponsibleStudentlUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getResponsibleStudentUserById(ctx *gin.Context) {
	var req getResponsibleStudentlUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	responsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, req.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	ctx.JSON(http.StatusOK, responsibleStudentUser)
}

type listResponsibleStudentUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAllResponsibleStudents(ctx *gin.Context) {
	var req listResponsibleStudentUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListResponsibleStudentParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	responsibleStudents, err := server.store.ListResponsibleStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, responsibleStudents)
}

type updateResponsibleStudentl struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	Cpf         int32     `json:"Cpf"`
	Gender      string    `json:"gender"`
}

func (server *Server) updateResponsibleStudentl(ctx *gin.Context) {
	var req updateResponsibleStudentl
	var reqID getResponsibleStudentlUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reponsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, reqID.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}
	arg := db.UpdateResponsibleStudentParams{
		ID:    reponsibleStudentUser.ID,
		Name:  req.Name,
		Email: req.Email,
	}
	responsiblestudentUpdated, err := server.store.UpdateResponsibleStudent(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not update this user")
		return
	}
	ctx.JSON(http.StatusOK, responsiblestudentUpdated)
}
