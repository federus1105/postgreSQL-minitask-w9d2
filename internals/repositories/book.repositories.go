package repositories

import (
	"context"
	"log"

	"github.com/federus1105/daysatu/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BooksRepository struct {
	db *pgxpool.Pool
}

func NewBooksRepository(db *pgxpool.Pool) *BooksRepository {
	return &BooksRepository{
		db: db,
	}
}
func (b *BooksRepository) GetBooksData(reqContext context.Context, offset, limit int) ([]models.Books, error) {
	sql := "SELECT id, title, author FROM books LIMIT $1 OFFSET $2"
	values := []any{limit, offset}
	rows, err := b.db.Query(reqContext, sql, values...)	
	if err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return []models.Books{}, err
	}
	defer rows.Close()
	var books []models.Books
	for rows.Next() {
		var book models.Books
			if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
				log.Println("Scan Error, ", err.Error())
				return []models.Books{}, err
			}
			books = append(books, book)
	}
	return books, nil
}
