package customers

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[customers.handler] req id :", r.Context().Value("request-id"))
	fmt.Fprintln(w, "All customer data will be served")
}
