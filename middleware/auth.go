package middleware

import (
	"trial/lib"

	"github.com/gin-gonic/gin"
)

func tokenFailed(ctx *gin.Context) {
	if e := recover(); e != nil {
		lib.HandlerUnauthorized(ctx, "Unautorization error")
		ctx.Abort()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer tokenFailed(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)

		if isValidated {
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			panic("Err: token Invalid")
		}
	}
}
