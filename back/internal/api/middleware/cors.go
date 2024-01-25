package middleware

import "github.com/gin-gonic/gin"

func CORS(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
