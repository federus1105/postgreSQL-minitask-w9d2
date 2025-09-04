package utils

import (
	"errors"

	"github.com/federus1105/daysatu/internals/models"
)

func ValidateBody(body models.Body) error {
	if body.Id <= 0 {
		return errors.New("ID tidak boleh di bawah 0")
	}
	if len(body.Message) < 8 {
		return errors.New("Panjang pesan harus lebih dari 8")
	}
	return nil
}
