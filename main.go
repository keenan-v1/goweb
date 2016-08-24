package main

import (
	"github.com/xorith/goweb/router"

	// Routes
	_ "github.com/xorith/goweb/routes/home"
	_ "github.com/xorith/goweb/routes/static"
)

func main() {
	router.LoadRoutes()
	router.Start()
}
