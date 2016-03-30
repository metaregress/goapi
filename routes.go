package main

import (
	"github.com/metaregress/api/actiongenerator"
	"github.com/metaregress/api/apiutil"
	"github.com/metaregress/api/todos"
)

func GetRoutes() apiutil.Routes {
	routes := apiutil.Routes{}
	todosRoutes := todos.GetRoutes()
	actiongenRoutes := actiongenerator.GetRoutes()
	routes = append(routes, todosRoutes...)
	routes = append(routes, actiongenRoutes...)
	routes = append(routes, apiutil.Route{
		"Index",
		"GET",
		"/",
		Index,
	})

	return routes
}
