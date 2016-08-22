package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xorith/server/routes"
)

// Router stores the router
// also comments are hard
type Router struct {
	router mux.Router
}

//RouteHandler is a wrapper to pass vars from the path
type RouteHandler func(http.ResponseWriter, *http.Request, map[string]string)

//ServeHTTP passes the vars along to the proper handler
func (rh RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rh(w, r, vars)
}

var (
	router = mux.NewRouter()
)

//LoadRoutes loads routes
func LoadRoutes() {
	for v := range routes.GetRoutes() {
		log.Printf("Loading %s (%s)", v.Name, v.Path)
		router.HandleFunc(v.Path, RouteHandler(v.Controller).ServeHTTP)
	}
}

//Start starts the router
func Start() {
	log.Println("STARTING SERVER")
	log.Fatal(http.ListenAndServe(":8000", router))
}
