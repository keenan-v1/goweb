package home

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

//TestHome is an example of a route test
func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := mux.Vars(req)
	HomeRoute.Controller(req, vars, HomeRoute.GetModel())
	mData := HomeRoute.GetModel().Data.(Model)
	if mData.Request == "" {
		t.Fail()
	}
	if mData.Vars == "" {
		t.Fail()
	}
	if mData.Route == "" {
		t.Fail()
	}
}
