package middleware

import (
	"back/internal/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authorization")

	var exist bool
	err := database.MySQL.Get(&exist, `SELECT EXISTS(SELECT * FROM user_sessions WHERE session_id = ?)`, authToken)
	if err != nil {
		fmt.Println(err)
	}

	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
