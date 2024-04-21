package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"myclass/token"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey   = "authorization"
	authorizationTypeBearer  = "bearer"
	authorizatiionPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("Unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accesToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accesToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		//session, _ := ctx.Get("session")
		//userSession := session.(*user.UserSession)
		//if userSession == nil {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User session not found"})
		//	return
		//}
		//if userSession.TokenAccess != accesToken {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token not accepted for this session"})
		//	return
		//}

		ctx.Set(authorizatiionPayloadKey, payload)
		ctx.Next()
	}
}
