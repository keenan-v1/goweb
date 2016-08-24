package static

import "github.com/xorith/goweb/routes"

//StaticHome is our route definition
var (
	StaticHome *routes.StaticRoute
)

// init() will define StaticHome and register it with routes.
// TODO: This should probably be a configuration, not hard-coded
func init() {
	StaticHome = &routes.StaticRoute{
		Name:     "staticHome",
		Path:     "/",
		Filepath: "./staticFiles",
	}
	routes.RegisterStatic(StaticHome)
}
