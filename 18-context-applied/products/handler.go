package products

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[products.handler] req id :", r.Context().Value("request-id"))
	ps := NewProductsService()
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(ps.GetAll(r.Context())); err != nil {
			http.Error(w, "error encoding data", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
		}
		ps.AddNew(r.Context(), newProduct)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
