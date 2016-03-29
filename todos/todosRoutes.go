package todos

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/todo",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},
	Route{
		"TodoShow",
		"DELETE",
		"/todo/{todoId}",
		TodoDelete,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todo",
		TodoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todo",
		TodoDelete,
	},
	Route{
		"TodoUpdate",
		"PATCH",
		"/todo/{todoId}",
		TodoUpdate,
	},
}

func GetRoutes() Routes {
	return routes
}
