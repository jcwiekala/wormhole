/*
 * Kyma Gateway Metadata API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"//",
		Index,
	},

	{
		"GetHealth",
		strings.ToUpper("Get"),
		"//v1/health",
		GetHealth,
	},

	{
		"PublishEvent",
		strings.ToUpper("Post"),
		"//v1/events",
		PublishEvent,
	},

	{
		"DeleteServiceByServiceId",
		strings.ToUpper("Delete"),
		"//v1/metadata/services/{serviceId}",
		DeleteServiceByServiceId,
	},

	{
		"GetServiceByServiceId",
		strings.ToUpper("Get"),
		"//v1/metadata/services/{serviceId}",
		GetServiceByServiceId,
	},

	{
		"GetServices",
		strings.ToUpper("Get"),
		"//v1/metadata/services",
		GetServices,
	},

	{
		"RegisterService",
		strings.ToUpper("Post"),
		"//v1/metadata/services",
		RegisterService,
	},

	{
		"UpdateService",
		strings.ToUpper("Put"),
		"//v1/metadata/services/{serviceId}",
		UpdateService,
	},
}