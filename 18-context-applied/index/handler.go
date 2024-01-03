package index

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[index.handler] req id :", r.Context().Value("request-id"))
	fmt.Fprintln(w, "Hello World!")
}
