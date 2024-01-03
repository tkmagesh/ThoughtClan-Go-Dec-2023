package main

import (
	"context-app/customers"
	"context-app/index"
	"context-app/products"
	"log"
	"net/http"
)

func main() {
	myServer := NewMyServer()
	myServer.UseMiddleware(LoggerMiddleware)
	myServer.UseMiddleware(TraceIdMiddleware)

	myServer.AddRoute("/", index.Handler)
	myServer.AddRoute("/customers", customers.Handler)
	myServer.AddRoute("/products", products.Handler)
	if err := http.ListenAndServe(":8080", myServer); err != nil {
		log.Println(err)
		return
	}
}
