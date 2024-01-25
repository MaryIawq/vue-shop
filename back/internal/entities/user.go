package entities

type User struct {
	ID                 int     `json:"id" db:"id"`
	Username           string  `json:"username" db:"username"`
	Email              string  `json:"email" db:"email"`
	Password           string  `json:"password" db:"password"`
	LastLoginTime      *int    `json:"last_login_time" db:"last_login_time"`
	LastLoginUserAgent *string `json:"last_login_user_agent" db:"last_login_user_agent"`
}
