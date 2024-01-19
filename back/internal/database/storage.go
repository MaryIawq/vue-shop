package database

import (
	"back/internal/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var MySQL *sqlx.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.JSON.MySQL.Username,
		config.JSON.MySQL.Password,
		config.JSON.MySQL.Host,
		config.JSON.MySQL.Port,
		config.JSON.MySQL.Database,
	)
	fmt.Println(dsn)
	var err error
	MySQL, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}
