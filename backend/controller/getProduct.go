package controller

import (
	"log"
	"net/http"
	"storage/models"
)

// DB-Verbindung (globaler Zugriff angenommen; anpassen, falls nötig)
func GetProducts(w http.ResponseWriter, r *http.Request) ([]models.Products, error) {

	rows, err := DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing products rows:", err)
		}
	}()

	var products []models.Products
	for rows.Next() {
		var product models.Products
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Amount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, err
}
