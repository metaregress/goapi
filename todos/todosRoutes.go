package todos

import (
	"github.com/metaregress/api/apiutil"
)

var routes = apiutil.Routes{
	apiutil.Route{
		"TodoIndex",
		"GET",
		"/todo",
		TodoIndex,
	},
	apiutil.Route{
		"TodoShow",
		"GET",
		"/todo/{todoId}",
		TodoShow,
	},
	apiutil.Route{
		"TodoShow",
		"DELETE",
		"/todo/{todoId}",
		TodoDelete,
	},
	apiutil.Route{
		"TodoCreate",
		"POST",
		"/todo",
		TodoCreate,
	},
	apiutil.Route{
		"TodoDelete",
		"DELETE",
		"/todo",
		TodoDelete,
	},
	apiutil.Route{
		"TodoUpdate",
		"PATCH",
		"/todo/{todoId}",
		TodoUpdate,
	},
}

func GetRoutes() apiutil.Routes {
	return routes
}
