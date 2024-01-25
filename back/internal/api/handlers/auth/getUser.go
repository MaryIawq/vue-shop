package authAPI

import (
	"back/internal/database"
	"back/internal/entities"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authorization")

	var user entities.User

	err := database.MySQL.Get(&user, `
	SELECT 
			u.id,
			u.username,
			u.email,
			u.password,
			u.last_login_time,
			u.last_login_user_agent
		FROM users AS u
		LEFT JOIN user_sessions AS us 
		ON u.id = us.user_id
		WHERE u.id = us.user_id AND us.session_id = ?;
	`,
		authToken)

	switch err {
	case sql.ErrNoRows:
		ctx.AbortWithStatus(http.StatusUnauthorized)
	case nil:
	default:
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, user)
}
