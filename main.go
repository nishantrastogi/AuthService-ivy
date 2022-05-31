package main

import (
	"authservice/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080"},
		AllowCredentials: true,
	})

	r := router.Router()

	handler := c.Handler(r)

	fmt.Println("Listening at http://localhost:8200")
	log.Fatal(http.ListenAndServe(":8200", handler))
}
