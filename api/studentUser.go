package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	db "myclass/db/sqlc"
	"myclass/token"
	"myclass/util"
	"net/http"
	"time"
)

type createStudentUserRequest struct {
	Username             string `json:"username" binding:"required,alphanum"`
	Password             string `json:"password" binding:"required,min=8,passwd"`
	Name                 string `json:"name" binding:"required"`
	DateOfBirth          string `json:"date_of_birth"`
	Gender               string `json:"gender" binding:"required,gender"`
	ResponsibleStudentID int64  `json:"responsible_student_id"`
}

type studentUserResponse struct {
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Gender            string    `json:"gender"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	CreatedAt         time.Time `json:"createdAt"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
}

func newStudentUserResponse(studentUser db.StudentUser) studentUserResponse {
	return studentUserResponse{
		Name:              studentUser.Name,
		Username:          studentUser.Username,
		Gender:            studentUser.Gender,
		DateOfBirth:       studentUser.DateOfBirth,
		CreatedAt:         studentUser.CreatedAt,
		PasswordChangedAt: studentUser.PasswordChangedAt,
	}
}

func (server *Server) createStudentUser(ctx *gin.Context) {
	var req createStudentUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqResponsible getResponsibleStudentlUserRequest
	if err := ctx.ShouldBindUri(&reqResponsible); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	responsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, reqResponsible.UserName)
	fmt.Println("reqResponsible", reqResponsible)
	fmt.Println("responsibleStudentUser", responsibleStudentUser)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This responsibleUser does not exists, please create it firstly")
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStudentUserParams{
		Username:             req.Username,
		HashedPassword:       hashedPassword,
		Name:                 req.Name,
		DateOfBirth:          dateOfBirth,
		Gender:               req.Gender,
		ResponsibleStudentID: responsibleStudentUser.ID,
	}

	studentUser, err := server.store.CreateStudentUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newStudentUserResponse(studentUser)

	ctx.JSON(http.StatusOK, response)
}

type getStudentlUserRequest struct {
	UserName string `uri:"username" binding:"required"`
}

type getResponsibleStudentlUserIDRequest struct {
	UserName string `uri:"responsibleStudent" binding:"required,min=1"`
}

func (server *Server) getStudentUserById(ctx *gin.Context) {
	var req getStudentlUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	studentUser, err := server.store.GetStudentUser(ctx, req.UserName)
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
	Username             string `json:"username"`
	Password             string `json:"password"`
	Name                 string `json:"name"`
	DateOfBirth          string `json:"date_of_birth"`
	Gender               string `json:"gender"`
	ResponsibleStudentID int64  `json:"responsible_student_id"`
}

func (server *Server) updateStudentlUser(ctx *gin.Context) {
	var req updateStudentlUser
	var reqID getStudentlUserRequest
	var reqResponsibleStudent getResponsibleStudentlUserIDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqResponsibleStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	studentUser, err := server.store.GetStudentUser(ctx, reqID.UserName)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
		return
	}

	responsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, reqResponsibleStudent.UserName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "This user does not have responsibleUser, please create it firstly")
		return
	}

	if responsibleStudentUser.ID != studentUser.ResponsibleStudentID {
		ctx.JSON(http.StatusNotFound, "The responsibleUser to this user is not correct, please verify it")
		return
	}

	arg := db.UpdateStudentUserParams{
		ID:                   studentUser.ID,
		Name:                 req.Name,
		Username:             req.Username,
		HashedPassword:       req.Password,
		ResponsibleStudentID: req.ResponsibleStudentID,
	}
	studentUserUpdated, err := server.store.UpdateStudentUser(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not update this user")
		return
	}

	authPayload := ctx.MustGet(authorizatiionPayloadKey).(*token.Payload)
	if studentUser.Username != authPayload.Username {
		err := errors.New("Account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, studentUserUpdated)
}

type loginStudentUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginStudentUserResponse struct {
	AccessToken string              `json:"access_token"`
	User        studentUserResponse `json:"user"`
}

func (server *Server) loginStudentUser(ctx *gin.Context) {
	var req loginStudentUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	studentUser, err := server.store.GetStudentUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, studentUser.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accesToken, err := server.tokenMaker.CreateToken(
		studentUser.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := loginStudentUserResponse{
		AccessToken: accesToken,
		User:        newStudentUserResponse(studentUser),
	}

	ctx.JSON(http.StatusOK, response)
}
