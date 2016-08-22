package main

import (
	"github.com/xorith/server/router"

	// Routes
	_ "github.com/xorith/server/routes/home"
)

func main() {
	router.LoadRoutes()
	router.Start()
}
