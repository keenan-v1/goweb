package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	routes map[string]*Route
)

func init() {
	routes = map[string]*Route{}
}

// Route is a path on the webserver
// TODO: User access levels
type Route struct {
	Path          string
	Controller    func(http.ResponseWriter, *http.Request, map[string]string)
	TemplateFiles []string
	View          *template.Template
	Name          string
	Title         string
}

//Register registers a route
func Register(route Route) *Route {
	routes[route.Name] = &route
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
	return &route
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
