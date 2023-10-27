package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"github.com/theluckiesthuman/yakshop/internal/adapter"
	"github.com/theluckiesthuman/yakshop/internal/config"
)

func main() {
	e := echo.New()
	//allow CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	e.Use(echo.WrapMiddleware(c.Handler))
	g := e.Group("/yak-shop")
	adapter.RegisterYakHandlers(g)
	port := fmt.Sprintf(":%s", config.Cfg.PORT)
	log.Printf("Serving http on %s\n", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to serve http server: %v", err)
	}
}
