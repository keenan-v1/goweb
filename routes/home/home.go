package home

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xorith/server/routes"
)

var (
	this *routes.Route
)

func init() {
	this = routes.Register(routes.Route{
		Name:          "home",
		Title:         "Home Page",
		Path:          "/",
		Controller:    Home,
		TemplateFiles: []string{"./routes/home/home.html"},
	})
}

// Home page handler (controller)
func Home(r http.ResponseWriter, req *http.Request, vars map[string]string) {
	s := fmt.Sprintf("%#v\n%#v\n%#v\n", this, req, vars)
	if this == nil || this.View == nil {
		r.Write([]byte(s))
		r.Write([]byte("<br><b>Some shit went wrong man...</b>"))
		return
	}
	this.View.Execute(r, nil)
	log.Println(s)
	r.Write([]byte(s))
}
