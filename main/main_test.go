package main

import (
	"strings"
	"testing"

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
	assert.Equal(t, "/api/users/:id\\:customAction", genericRouteCustomAction)

	router.Find("POST", "/api/users/123:otherAction", echoContext)
	genericRouteOtherAction := echoContext.Path()
	assert.Equal(t, "/api/users/\\:id:otherAction", genericRouteOtherAction)
}

func void(c echo.Context) error {
	return nil
}
