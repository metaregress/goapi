package main

import (
	"github.com/metaregress/api/todos"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = todos.GetRoutes()

// routes.append(Route{
// 		"Index",
// 		"GET",
// 		"/",
// 		Index,
// 	},)
