package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"storage/database"
	"storage/models"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) error {
	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Lesen des Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	//Todo: Unmarshaling der JSON-Daten - Unmarshal dont work #fix
	var products []models.Products //Produkte ist ein Slice von Produktstrukturen
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Json", products)
		return err
	}

	// Einfügen der JSON-Daten in die Datenbank
	stmt, err := db.Prepare("INSERT INTO products (ID, NAME, DESCRIPTION, PRICE, AMOUNT ...) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		// Fehlerbehandlung
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Schließen des Statements verzögert (nach Abschluss der Funktion)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			// Fehlerbehandlung
		}
	}(stmt)

	// Wenn nur ein Produkt vorhanden ist
	if len(products) == 1 {
		product := products[0]
		values := []interface{}{product.ID, product.Name, product.Description, product.Price, product.Amount}
		_, err := stmt.Exec(values...)
		if err != nil {
			// Fehlerbehandlung
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	} else {
		// Wenn mehrere Produkte vorhanden sind
		for _, product := range products {
			values := []interface{}{product.ID, product.Name, product.Description, product.Price, product.Amount}
			_, err := stmt.Exec(values...)
			if err != nil {
				// Fehlerbehandlung
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}
		}

	}

	w.WriteHeader(http.StatusCreated) // Erfolgsmeldung senden
	return err

}
