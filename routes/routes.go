package routes

import "net/http"

var (
	routes map[string]Route
)

func init() {
	routes = map[string]Route{}
}

// Route defines various endpoints and routes for the server
type Route interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	GetName() string
	GetPath() string
	GetTitle() string
}

//GetRoute gets a single route by-name
func GetRoute(name string) Route {
	if v, b := routes[name]; b {
		return v
	}
	return nil
}

//GetRoutes returns an channel to iterate through all registered routes
func GetRoutes() chan Route {
	ch := make(chan Route)
	go func() {
		for _, v := range routes {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
