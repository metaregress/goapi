package main

import (
	"github.com/metaregress/api/apiutil"
	"github.com/metaregress/api/todos"
)

func GetRoutes() apiutil.Routes {
	var routes = todos.GetRoutes()
	routes = append(routes, apiutil.Route{
		"Index",
		"GET",
		"/",
		Index,
	})

	return routes
}
