package middleware

import "github.com/gin-gonic/gin"

func Cors(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}
	ctx.Next()
}