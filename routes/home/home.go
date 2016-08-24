package home

import (
	"fmt"
	"net/http"

	"github.com/xorith/goweb/routes"
)

//HomeRoute is our route definition
var (
	HomeRoute *routes.TemplateRoute
)

// init() will define HomeRoute and register it with routes.
func init() {
	HomeRoute = &routes.TemplateRoute{
		Name:          "home",
		Title:         "Home Page",
		Path:          "/",
		Controller:    Home,
		TemplateFiles: []string{"./routes/home/home.html"},
	}
	routes.RegisterTemplate(HomeRoute)
}

// Model is the model home. hahaha no.
type Model struct {
	Request string
	Vars    string
	Route   string
}

// Home page handler (controller)
func Home(req *http.Request, vars map[string]string, model *routes.Model) {
	mData := Model{}
	mData.Request = fmt.Sprintf("%#v", req)
	mData.Vars = fmt.Sprintf("%#v", vars)
	mData.Route = fmt.Sprintf("%#v", HomeRoute)
	model.Data = mData
}
