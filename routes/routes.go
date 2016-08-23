package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

var (
	routes map[string]*Route
)

func init() {
	routes = map[string]*Route{}
}

// Route is a path on the webserver
// TODO: User access levels
// TODO: TODO: Don't be so stupid
// TODO: TODO: TODO: You're not my real mom.
type Route struct {
	Path          string
	Controller    func(*http.Request, map[string]string, *Model)
	TemplateFiles []string
	View          *template.Template
	Name          string
	Title         string
	model         *Model
}

// Model is the base model for a route
type Model struct {
	PageTitle string
	Data      interface{}
}

// GetModel returns the model for this route
func (route *Route) GetModel() *Model {
	if route.model == nil {
		route.model = &Model{PageTitle: route.Title}
	}
	return route.model
}

//ServeHTTP passes the vars along to the proper handler
func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route.Controller(r, vars, route.GetModel())
	route.View.Execute(w, route.GetModel())
}

//Register registers a route
func Register(route *Route) {
	routes[route.Name] = route
	var paths []string
	for _, v := range route.TemplateFiles {
		if path, err := filepath.Abs(v); err != nil {
			log.Printf("Error loading %s\n%#v", v, err.Error())
		} else {
			paths = append(paths, path)
		}
	}
	var err error
	if route.View, err = template.ParseFiles(paths...); err != nil {
		log.Printf("Error parsing templates: %#v", err.Error())
	}
}

//GetRoute gets a single route by-name
func GetRoute(name string) *Route {
	if v, b := routes[name]; b {
		return v
	}
	return nil
}

//GetRoutes returns an channel to iterate through all registered routes
func GetRoutes() chan *Route {
	ch := make(chan *Route)
	go func() {
		for _, v := range routes {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
