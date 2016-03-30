package actiongenerator

import (
	"github.com/metaregress/api/apiutil"
)

var routes = apiutil.Routes{
	apiutil.Route{
		"Randomize",
		"GET",
		"/actiongenerator/randomize",
		Randomize,
	},
}

func GetRoutes() apiutil.Routes {
	return routes
}
