package api

// import (
//
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/mvrilo/go-cpf"
//	db "myclass/db/sqlc"
//	"myclass/token"
//	"myclass/util"
//	"net/http"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	_ "github.com/lib/pq"
//
// )
//
//	type createProfessionalUserRequest struct {
//		Name           string `json:"name" binding:"required"`
//		Username       string `json:"username" binding:"required,alphanum"`
//		Password       string `json:"password" binding:"required,min=8,passwd"`
//		Gender         string `json:"gender" binding:"required,gender"`
//		Email          string `json:"email" binding:"required,email"`
//		DateOfBirth    string `json:"date_of_birth"`
//		Cpf            string `json:"cpf"`
//		ClassHourPrice string `json:"class_hour_price"`
//		ImageID        int64  `json:"image_id"`
//	}
//
//	type professionalUserResponse struct {
//		Name              string    `json:"name"`
//		Username          string    `json:"username"`
//		Gender            string    `json:"gender"`
//		Email             string    `json:"email"`
//		DateOfBirth       time.Time `json:"date_of_birth"`
//		Cpf               string    `json:"cpf"`
//		ClassHourPrice    string    `json:"class_hour_price"`
//		ImageID           int64     `json:"image_id"`
//		CreatedAt         time.Time `json:"createdAt"`
//		PasswordChangedAt time.Time `json:"passwordChangedAt"`
//	}
//
//	func newProfessionalUserResponse(professionalUser db.ProfessionalUser) professionalUserResponse {
//		return professionalUserResponse{
//			Name:              professionalUser.Name,
//			Username:          professionalUser.Username,
//			Gender:            professionalUser.Gender,
//			Email:             professionalUser.Email,
//			DateOfBirth:       professionalUser.DateOfBirth,
//			Cpf:               professionalUser.Cpf,
//			ClassHourPrice:    professionalUser.ClassHourPrice,
//			ImageID:           professionalUser.ImageID,
//			CreatedAt:         professionalUser.CreatedAt,
//			PasswordChangedAt: professionalUser.PasswordChangedAt,
//		}
//	}
//
//	func (server *Server) createProfessionalUser(ctx *gin.Context) {
//		var req createProfessionalUserRequest
//		if err := ctx.ShouldBindJSON(&req); err != nil {
//			ctx.JSON(http.StatusBadRequest, errorResponse(err))
//			return
//		}
//
//		hashedPassword, err := util.HashPassword(req.Password)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//			return
//		}
//
//		dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, errorResponse(err))
//			return
//		}
//
//		if validCpf, err := cpf.Valid(req.Cpf); !validCpf {
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//				return
//			}
//			ctx.JSON(http.StatusBadRequest, fmt.Errorf("Invalid cpf"))
//			return
//		}
//
//		arg := db.CreateProfessionalUserParams{
//			Name:           req.Name,
//			Username:       req.Username,
//			HashedPassword: hashedPassword,
//			Gender:         req.Gender,
//			Email:          req.Email,
//			DateOfBirth:    dateOfBirth,
//			Cpf:            req.Cpf,
//			ImageID:        req.ImageID,
//			UpdatedAt:      time.Now(),
//			ClassHourPrice: req.ClassHourPrice,
//		}
//
//		professionalUser, err := server.store.CreateProfessionalUser(ctx, arg)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//			return
//		}
//
//		response := newProfessionalUserResponse(professionalUser)
//
//		ctx.JSON(http.StatusOK, response)
//	}
//
//	type getProfissionalUserRequest struct {
//		UserName string `uri:"username" binding:"required"`
//	}
//
//	func (server *Server) getProfessionalUserById(ctx *gin.Context) {
//		var req getProfissionalUserRequest
//		if err := ctx.ShouldBindUri(&req); err != nil {
//			ctx.JSON(http.StatusBadRequest, errorResponse(err))
//			return
//		}
//		professionalUser, err := server.store.GetProfessionalUser(ctx, req.UserName)
//		if err != nil {
//			fmt.Println(err)
//			ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
//			return
//		}
//		ctx.JSON(http.StatusOK, professionalUser)
//	}
//
//	type listProfessionalUsersRequest struct {
//		PageID   int32 `form:"page_id" binding:"required,min=1"`
//		PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
//	}
//
//	func (server *Server) listAllProfessionalUsers(ctx *gin.Context) {
//		var req listProfessionalUsersRequest
//		if err := ctx.ShouldBindQuery(&req); err != nil {
//			ctx.JSON(http.StatusBadRequest, errorResponse(err))
//			return
//		}
//
//		arg := db.ListProfessionalUserParams{
//			Limit:  req.PageSize,
//			Offset: (req.PageID - 1) * req.PageSize,
//		}
//
//		professionalUsers, err := server.store.ListProfessionalUser(ctx, arg)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusOK, professionalUsers)
//	}
//
//	type updateProfessionalUser struct {
//		Name           string `json:"name"`
//		Username       string `json:"username"`
//		Password       string `json:"password"`
//		Email          string `json:"email"`
//		DateOfBirth    string `json:"date_of_birth"`
//		ClassHourPrice string `json:"class_hour_price"`
//	}
//func (server *Server) updateProfessionalUser(ctx *gin.Context) {
//	var req updateProfessionalUser
//	var reqID getProfissionalUserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	if err := ctx.ShouldBindUri(&reqID); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	professionalUser, err := server.store.GetProfessionalUser(ctx, reqID.UserName)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusNotFound, "This user does not exists, please create it firstly")
//		return
//	}
//
//	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.UpdateProfessionalUserParams{
//		ID:             professionalUser.ID,
//		Name:           req.Name,
//		Username:       req.Username,
//		HashedPassword: req.Password,
//		Email:          req.Email,
//		DateOfBirth:    dateOfBirth,
//		ClassHourPrice: req.ClassHourPrice,
//	}
//	professionalUserUpdated, err := server.store.UpdateProfessionalUser(ctx, arg)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusBadRequest, "Could not update this user")
//		return
//	}
//
//	authPayload := ctx.MustGet(authorizatiionPayloadKey).(*token.Payload)
//	if professionalUser.Username != authPayload.Username {
//		err := errors.New("Account does not belong to the authenticated user")
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, professionalUserUpdated)
//}

//
//func (server *Server) deleteProfessionalUser(ctx *gin.Context) {
//	var reqID getProfissionalUserRequest
//	if err := ctx.ShouldBindUri(&reqID); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//	professionalUser, err := server.store.GetProfessionalUser(ctx, reqID.UserName)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusNotFound, "This user does not exist")
//		return
//	}
//
//	authPayload := ctx.MustGet(authorizatiionPayloadKey).(*token.Payload)
//	if professionalUser.Username != authPayload.Username {
//		err := errors.New("Account does not belong to the authenticated user")
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	_, err = server.store.DeleteProfessionalUser(ctx, professionalUser.ID)
//
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(http.StatusBadRequest, "Could not delete this user")
//		return
//	}
//}
//
//type loginProfessionalUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
//type loginProfessionalUserResponse struct {
//	AccessToken string                   `json:"access_token"`
//	User        professionalUserResponse `json:"user"`
//}
//
//func (server *Server) loginProfessionalUser(ctx *gin.Context) {
//	var req loginProfessionalUserRequest
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	professionalUser, err := server.store.GetProfessionalUser(ctx, req.Username)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	err = util.CheckPassword(req.Password, professionalUser.HashedPassword)
//	if err != nil {
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	accesToken, err := server.tokenMaker.CreateToken(
//		professionalUser.Username,
//		server.config.AccessTokenDuration,
//	)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	response := loginProfessionalUserResponse{
//		AccessToken: accesToken,
//		User:        newProfessionalUserResponse(professionalUser),
//	}
//
//	ctx.JSON(http.StatusOK, response)
//}
