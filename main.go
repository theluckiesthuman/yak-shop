package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/yakshop/internal/adapter"
	"github.com/theluckiesthuman/yakshop/internal/config"
)

func main() {
	e := echo.New()
	g := e.Group("/yak-shop")
	adapter.RegisterYakHandlers(g)
	
	port := fmt.Sprintf(":%s", config.Cfg.Port)
	log.Printf("Serving http on %s\n", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to serve http server: %v", err)
	}
}
