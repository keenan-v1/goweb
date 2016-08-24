package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// TemplateRoute is a route that uses a template in an MVC design
type TemplateRoute struct {
	Path          string
	Controller    func(*http.Request, map[string]string, *Model)
	TemplateFiles []string
	View          *template.Template
	Name          string
	Title         string
	model         *Model
}

// Model is the base model for a TemplateRoute
type Model struct {
	PageTitle string
	Data      interface{}
}

// GetModel returns the model for this route
func (route *TemplateRoute) GetModel() *Model {
	if route.model == nil {
		route.model = &Model{PageTitle: route.GetTitle()}
	}
	return route.model
}

// GetName returns the name of the route
func (route *TemplateRoute) GetName() string {
	return route.Name
}

// GetPath returns the server path of the route
func (route *TemplateRoute) GetPath() string {
	return route.Path
}

// GetTitle returns the page title of the route
func (route *TemplateRoute) GetTitle() string {
	return route.Title
}

//ServeHTTP passes the vars along to the proper handler
func (route *TemplateRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route.Controller(r, vars, route.GetModel())
	route.View.Execute(w, route.GetModel())
}

//RegisterTemplate registers a templated route
func RegisterTemplate(route *TemplateRoute) {
	routes[route.GetName()] = route
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
