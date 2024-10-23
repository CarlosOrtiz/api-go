package utils

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type PaginateParams struct {
	Page     int
	Quantity int
}

func Paginate(w http.ResponseWriter, r *http.Request, db *gorm.DB) (*gorm.DB, int, int, error) {
	pageStr := r.URL.Query().Get("page")
	quantityStr := r.URL.Query().Get("quantity")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		quantity = 10
	}

	offset := (page - 1) * quantity

	return db.Offset(offset).Limit(quantity), page, quantity, nil
}
