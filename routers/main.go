package routers

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	RoutersAuth(r.Group("/auth"))
}
