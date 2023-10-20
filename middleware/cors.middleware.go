package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Signature, TimeStamp")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Accept, Content-Type, Authorization, TimeStamp, Signature")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}
	ctx.Next()
}