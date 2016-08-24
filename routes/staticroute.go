package routes

// StaticRoute is a route that points to a specific folder
type StaticRoute struct {
	Filepath string
	Path     string
	Name     string
}

// GetName returns the name of the route
func (route *StaticRoute) GetName() string {
	return route.Name
}

// GetType returns Static
func (route *StaticRoute) GetType() RouteType {
	return Static
}

// GetPath returns the URL path for the route
func (route *StaticRoute) GetPath() string {
	return route.Path
}

// GetFilepath returns the filepath of the static route
func (route *StaticRoute) GetFilepath() string {
	return route.Filepath
}

//RegisterStatic registers a templated route
func RegisterStatic(route *StaticRoute) {
	routes[route.GetName()] = route
}
