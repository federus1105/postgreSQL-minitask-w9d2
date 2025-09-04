package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/federus1105/daysatu/internals/configs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// inisialisai
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env\nCause:", err.Error())
		return
	}
	log.Println(os.Getenv("DBUSER"))

	// inisialisasi db
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Ping to DB failed\nCause:", err.Error())
		return
	}
	log.Println("DB Connected")

	rounter := gin.Default()
	rounter.Use(MyLogger)
	rounter.Use(CORSMiddleware)

	rounter.GET("/ping", func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-ID")
		contentType := ctx.GetHeader("Content-Type")
		ctx.JSON(http.StatusOK, gin.H{
			"message":     "pong",
			"requestId":   requestId,
			"contentType": contentType,
		})
	})
	rounter.GET("/ping/:id/:param2", func(ctx *gin.Context) {
		pingId := ctx.Param("id")
		param2 := ctx.Param("param2")
		query := ctx.Query("query")
		ctx.JSON(http.StatusOK, gin.H{
			"param":  pingId,
			"param2": param2,
			"query":  query,
		})
	})
	rounter.GET("/books", func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 5
		offset := (page - 1) * limit
		sql := "SELECT id, title, author FROM books LIMIT $1 OFFSET $2"
		values := []any{limit, offset}
		rows, err := db.Query(ctx.Request.Context(), sql, values...)
		if err != nil {
			log.Println("Internal server Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"succes": false,
				"data":   []any{},
			})
			return
		}
		defer rows.Close()
		var books []Books
		for rows.Next() {
			var book Books
			if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
				log.Println("Scan Error, ", err.Error())
				return
			}
			books = append(books, book)
		}
		if len(books) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"succes": false,
				"data":   []any{},
				"page":   page,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"succes": true,
			"data":   books,
		})
	})
	rounter.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Message: "Rute salah",
			Status:  "Rute tidak ditemukan",
		})
	})

	rounter.POST("/products", func(ctx *gin.Context) {
		var body Product
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
		values := []any{body.Name, body.PromoId, body.Price}
		var newProduct Product
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
			log.Println("Internal server Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"succes": false,
				"data":   []any{},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"succes": true,
			"data":   newProduct,
		})
	})
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
	rounter.Run("localhost:8080")
}

type Product struct {
	Name       string    `json:"name" `
	PromoId    *int      `json:"promo_id"`
	Price      int       `json:"price,omitempty"`
	Updatet_at time.Time `json:"update"`
	Id         int       `json:"id,omitempty"`
}

type Books struct {
	Id     int    `db:"id" json:"id"`
	Title  string `db:"title" json:"title_buku"`
	Author string `db:"author" json:"pembuat"`
}

func MyLogger(ctx *gin.Context) {
	log.Println("Start")
	start := time.Now()
	ctx.Next() // next digunakan untuk lanjut ke middleware selanjutnya
	duration := time.Since(start)
	log.Printf("Durasi Request: %dus\n", duration.Microseconds())
}

func CORSMiddleware(ctx *gin.Context) {
	whitelist := []string{"http://localhost:5500", "http://127.0.0.1:3001"}
	origin := ctx.GetHeader("Origin")
	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}

	ctx.Header("Access-Control-Allow-Methods", "GET")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}

type Response struct {
	Message string
	Status  string
}
