package repositories

import (
	"context"
	"log"
	"time"

	"github.com/federus1105/daysatu/internals/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
func (p *ProductRepository) AddNewProduct(rctx context.Context, body models.Product) (models.Product, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
	values := []any{body.Name, body.PromoId, body.Price}
	var newProduct models.Product
	if err := p.db.QueryRow(rctx, sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
		log.Println("Internal server Error: ", err.Error())
		return models.Product{}, err
	}
	return newProduct, nil
}
func (p *ProductRepository) InsertNewProduct(rctx context.Context, body models.Product) (pgconn.CommandTag, error) {
	sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3)"
	values := []any{body.Name, body.PromoId, body.Price}
	return p.db.Exec(rctx, sql, values...)
}

func (p *ProductRepository) UpdateDataProduct(rctx context.Context, body models.Product) (models.Product, error) {
	ProductiD := rctx.Param("id")
	sql := "UPDATE products SET name = $1, updatet_at = $2 WHERE id = $3 RETURNING id, name, updatet_at"
	values := []any{body.Name, time.Now(), ProductiD}
	var updatedProduct models.Product
	if err := p.db.QueryRow(rctx, sql, values...).Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Updatet_at); err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return models.Product{}, err
	}
	return updatedProduct, nil
}
