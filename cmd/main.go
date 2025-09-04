package main

import (
	"log"
	"net/http"
	"time"

	"github.com/federus1105/daysatu/internals/configs"
	"github.com/federus1105/daysatu/internals/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	// inisialisai
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("Failed to load env\nCause:", err.Error())
	// 	return
	// }
	// log.Println(os.Getenv("DBUSER"))

	// inisialisasi db
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}

	defer db.Close()
	if err := configs.TestDB(db); err != nil {
		log.Println("Ping to DB failed\nCause:", err.Error())
		return
	}
	log.Println("DB Connected")

	router := routers.InitRouter(db)

	rounter := gin.Default()

	rounter.PATCH("products/:id", func(ctx *gin.Context) {
		ProductiD := ctx.Param("id")
		var body Product
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		sql := "UPDATE products SET name = $1, updatet_at = $2 WHERE id = $3 RETURNING id, name, updatet_at"
		values := []any{body.Name, time.Now(), ProductiD}
		var updatedProduct Product
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Updatet_at); err != nil {
			log.Println("Internal Server Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"data":    []any{},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    updatedProduct,
		})
	})
	router.Run("localhost:8080")
}

type Product struct {
	Name       string    `json:"name" binding:"required" `
	PromoId    *int      `json:"promo_id"`
	Price      int       `json:"price,omitempty" binding:"required"`
	Updatet_at time.Time `json:"update"`
	Id         int       `json:"id,omitempty"`
}
