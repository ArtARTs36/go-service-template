package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func CORS(handler http.Handler) http.Handler {
	handleCORS := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods: []string{
			"HEAD",
			"GET",
			"POST",
			"PATCH",
		},
	}).Handler

	return handleCORS(handler)
}
