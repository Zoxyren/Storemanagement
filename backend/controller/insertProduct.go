package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"storage/models"
)

var DB sql.DB

func InsertProduct(w http.ResponseWriter, r *http.Request) error {
	// Lesen des Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var products *[]models.Products //Produkte ist ein Slice von Produktstrukturen
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("The following Json File can not be unmarshald", products)
		return err
	}
	// Einfügen der JSON-Daten in die Datenbank
	stmt, err := DB.Prepare("INSERT INTO products (id, name, description, price, amount) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		// Fehlerbehandlung
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Schließen des Statements verzögert (nach Abschluss der Funktion)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			slog.Error("Failed to close statement", err)
		}
	}(stmt)

	// Wenn nur ein Produkt vorhanden ist
	if len(*products) == 1 {
		product := (*products)[0]
		values := []interface{}{product.ID, product.Name, product.Description, product.Price, product.Amount}
		_, err := stmt.Exec(values...)
		if err != nil {
			// Fehlerbehandlung
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	} else {
		// Wenn mehrere Produkte vorhanden sind
		for _, product := range *products {
			values := []interface{}{product.ID, product.Name, product.Description, product.Price, product.Amount}
			_, err := stmt.Exec(values...)
			if err != nil {
				// Fehlerbehandlung
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}
		}
		defer func(DB *sql.DB) {
			err := DB.Close()
			if err != nil {
				slog.Error("Failed to close DB", err)
			}
		}(&DB)

	}

	w.WriteHeader(http.StatusCreated) // Erfolgsmeldung senden
	return err

}
