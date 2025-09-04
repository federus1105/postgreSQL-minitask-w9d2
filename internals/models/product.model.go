package models

import "time"

type Product struct {
	Name       string    `json:"name" binding:"required" `
	PromoId    *int      `json:"promo_id"`
	Price      int       `json:"price,omitempty" binding:"required"`
	Updatet_at time.Time `json:"update"`
	Id         int       `json:"id,omitempty"`
}
type ProductResponse struct{}
type NewProduct struct{}