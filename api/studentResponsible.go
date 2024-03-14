package api

//
//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/mvrilo/go-cpf"
//	db "myclass/db/sqlc"
//	"myclass/token"
//	"myclass/util"
//	"net/http"
//	"time"
//)
//
//type createStudentResponsbileRequest struct {
//	Name        string `json:"name" binding:"required"`
//	Username    string `json:"username" binding:"required,alphanum"`
//	DateOfBirth string `json:"date_of_birth"`
//	Email       string `json:"email" binding:"required,email"`
//	Cpf         string `json:"Cpf"`
//	Gender      string `json:"gender" binding:"required,gender"`
//	Password    string `json:"password" binding:"required,min=8,passwd"`
//}
//
//type responsibleStudentUserResponse struct {
//	Name              string    `json:"name"`
//	Username          string    `json:"username"`
//	DateOfBirth       time.Time `json:"date_of_birth"`
//	CreatedAt         time.Time `json:"createdAt"`
//	PasswordChangedAt time.Time `json:"passwordChangedAt"`
//}
//
//func newResponsibleStudentUserResponse(responsbibleStudentUser db.ResponsibleStudent) responsibleStudentUserResponse {
//	return responsibleStudentUserResponse{
//		Name:              responsbibleStudentUser.Name,
//		Username:          responsbibleStudentUser.Username,
//		DateOfBirth:       responsbibleStudentUser.DateOfBirth,
//		CreatedAt:         responsbibleStudentUser.CreatedAt,
//		PasswordChangedAt: responsbibleStudentUser.PasswordChangedAt,
//	}
//}
//
//func (server *Server) createResponsibleStudent(ctx *gin.Context) {
//	var req createStudentResponsbileRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	hashedPassword, err := util.HashPassword(req.Password)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	if validCpf, err := cpf.Valid(req.Cpf); !validCpf {
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusBadRequest, fmt.Errorf("Invalid cpf"))
//		return
//	}
//
//	arg := db.CreateResponsibleStudentParams{
//		Name:           req.Name,
//		HashedPassword: hashedPassword,
//		Username:       req.Username,
//		Gender:         req.Gender,
//		Email:          req.Email,
//		DateOfBirth:    dateOfBirth,
//		Cpf:            req.Cpf,
//	}
//
//	responsibleStudent, err := server.store.CreateResponsibleStudent(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, responsibleStudent)
//}
//
//type getResponsibleStudentlUserRequest struct {
//	UserName string `uri:"username" binding:"required"`
//}
//
//func (server *Server) getResponsibleStudentUserById(ctx *gin.Context) {
//	var req getResponsibleStudentlUserRequest
//	if err := ctx.ShouldBindUri(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	responsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, req.UserName)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
//		return
//	}
//	ctx.JSON(http.StatusOK, responsibleStudentUser)
//}
//
//type listResponsibleStudentUsersRequest struct {
//	PageID   int32 `form:"page_id" binding:"required,min=1"`
//	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
//}
//
//func (server *Server) listAllResponsibleStudents(ctx *gin.Context) {
//	var req listResponsibleStudentUsersRequest
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.ListResponsibleStudentParams{
//		Limit:  req.PageSize,
//		Offset: (req.PageID - 1) * req.PageSize,
//	}
//
//	responsibleStudents, err := server.store.ListResponsibleStudent(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//	ctx.JSON(http.StatusOK, responsibleStudents)
//}
//
//type updateResponsibleStudentl struct {
//	Name        string `json:"name"`
//	DateOfBirth string `json:"date_of_birth"`
//	Email       string `json:"email"`
//	Cpf         int32  `json:"cpf"`
//	Gender      string `json:"gender"`
//}
//
//func (server *Server) updateResponsibleStudent(ctx *gin.Context) {
//	var req updateResponsibleStudentl
//	var reqID getResponsibleStudentlUserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	if err := ctx.ShouldBindUri(&reqID); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	reponsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, reqID.UserName)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
//		return
//	}
//
//	arg := db.UpdateResponsibleStudentParams{
//		ID:    reponsibleStudentUser.ID,
//		Name:  req.Name,
//		Email: req.Email,
//	}
//	responsiblestudentUpdated, err := server.store.UpdateResponsibleStudent(ctx, arg)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusBadRequest, "Could not update this user")
//		return
//	}
//
//	authPayload := ctx.MustGet(authorizatiionPayloadKey).(*token.Payload)
//
//	if reponsibleStudentUser.Username != authPayload.Username {
//		err := errors.New("Account does not belong to the authenticated user")
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, responsiblestudentUpdated)
//}
//
//type loginResponsbibleStudentUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
//type loginResponsbibleStudentResponse struct {
//	AccessToken string                         `json:"access_token"`
//	User        responsibleStudentUserResponse `json:"user"`
//}
//
//func (server *Server) loginResponsibleStudent(ctx *gin.Context) {
//	var req loginResponsbibleStudentUserRequest
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	responsibleStudentUser, err := server.store.GetResponsibleStudent(ctx, req.Username)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	err = util.CheckPassword(req.Password, responsibleStudentUser.HashedPassword)
//	if err != nil {
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	accesToken, err := server.tokenMaker.CreateToken(
//		responsibleStudentUser.Username,
//		server.config.AccessTokenDuration,
//	)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	response := loginResponsbibleStudentResponse{
//		AccessToken: accesToken,
//		User:        newResponsibleStudentUserResponse(responsibleStudentUser),
//	}
//
//	ctx.JSON(http.StatusOK, response)
//}
