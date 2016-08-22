package main

import (
	"github.com/xorith/goweb/router"

	// Routes
	_ "github.com/xorith/goweb/routes/home"
)

func main() {
	router.LoadRoutes()
	router.Start()
}
