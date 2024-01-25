package productsAPI

import (
	"back/internal/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int     `json:"id" db:"id,int,NO"`
	Category     string  `json:"category" db:"category,varchar(50),NO"`
	ImageURL     string  `json:"imageUrl" db:"imageUrl,text,NO"`
	Title        string  `json:"title" db:"title,varchar(100),NO"`
	Price        float64 `json:"price" db:"price,float,NO"`
	Weight       float64 `json:"weight,omitempty" db:"weight,float,YES"`
	IsSugarFree  bool    `json:"is_sugar_free,omitempty" db:"is_sugar_free,tinyint(1),YES"`
	Protein      float64 `json:"protein,omitempty" db:"protein,float,YES"`
	Fat          float64 `json:"fat,omitempty" db:"fat,float,YES"`
	Carbs        float64 `json:"carbs,omitempty" db:"carbs,float,YES"`
	Calories     float64 `json:"calories,omitempty" db:"calories,float,YES"`
	InCart       bool    `json:"in_cart" db:"in_cart,tinyint(1),NO"`
	InFavourites bool    `json:"in_favourites" db:"in_favourites,tinyint(1),NO"`
}

func GetAll(ctx *gin.Context) {
	var products []Product
	err := database.MySQL.Select(&products, "SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, &products)
}
