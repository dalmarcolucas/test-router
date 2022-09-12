package main

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	e := echo.New()

	router := echo.NewRouter(e)

	router.Add(strings.ToUpper("PUT"), "/api/users/:id", void)
	router.Add(strings.ToUpper("POST"), "/api/users/:id\\:customAction", void)
	router.Add(strings.ToUpper("POST"), "/api/users/:id\\:otherAction", void)

	echoContext := e.NewContext(nil, nil)

	router.Find("GET", "/api/users/123", echoContext)
	genericRoute := echoContext.Path()
	assert.Equal(t, "/api/users/:id", genericRoute)

	router.Find("POST", "/api/users/123:customAction", echoContext)
	genericRouteCustomAction := echoContext.Path()
	assert.Equal(t, "/api/users/:id:customAction", genericRouteCustomAction)

	router.Find("POST", "/api/users/123:otherAction", echoContext)
	genericRouteOtherAction := echoContext.Path()
	assert.Equal(t, "/api/users/:id:otherAction", genericRouteOtherAction)
}

func TestMux(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/{id}", nil).Methods("GET")
	r.HandleFunc("/api/users/{id}:customAction", nil).Methods("POST")
	r.HandleFunc("/api/users/{id}:otherAction", nil).Methods("POST")

	routeMatch := mux.RouteMatch{}
	routeMatch2 := mux.RouteMatch{}
	routeMatch3 := mux.RouteMatch{}

	url1 := url.URL{Path: "/api/users/123"}
	request := http.Request{Method: "GET", URL: &url1}
	x := r.Match(&request, &routeMatch)

	assert.True(t, x)
	template, _ := routeMatch.Route.GetPathTemplate()
	assert.Equal(t, "/api/users/:id", template)

	url2 := url.URL{Path: "/api/users/123:customAction"}
	request2 := http.Request{Method: "POST", URL: &url2}
	x = r.Match(&request2, &routeMatch2)

	assert.True(t, x)
	template2, _ := routeMatch2.Route.GetPathTemplate()
	assert.Equal(t, "/api/users/{id}:customAction", template2)

	url3 := url.URL{Path: "/api/users/123:otherAction"}
	request3 := http.Request{Method: "POST", URL: &url3}
	x = r.Match(&request3, &routeMatch3)

	assert.True(t, x)
	template3, _ := routeMatch3.Route.GetPathTemplate()
	assert.Equal(t, "/api/users/{id}:otherAction", template3)

}

func void(c echo.Context) error {
	return nil
}
