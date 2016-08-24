package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xorith/goweb/routes"
)

// Router stores the router
// also comments are hard
type Router struct {
	router mux.Router
}

var (
	router = mux.NewRouter()
)

//LoadRoutes loads routes
func LoadRoutes() {
	for v := range routes.GetRoutes() {
		log.Printf("Loading %s (%s)", v.GetName(), v.GetPath())
		router.HandleFunc(v.GetPath(), v.ServeHTTP)
	}
}

//Start starts the router
func Start() {
	log.Println("STARTING SERVER")
	log.Fatal(http.ListenAndServe(":8000", router))
}
