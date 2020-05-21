package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/InviewTeam/BEST_Final/api"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
			AllowCredentials: true,
			Debug:            true,
		})

		h := c.Handler(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(h)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateEquipment",
		strings.ToUpper("Post"),
		"/add",
		api.Create,
	},

	Route{
		"DeleteEquipment",
		strings.ToUpper("Delete"),
		"/delete/{id}",
		api.Delete,
	},

	Route{
		"GetEquipments",
		strings.ToUpper("Get"),
		"/equipments",
		api.Get,
	},

	Route{
		"Login",
		strings.ToUpper("Post"),
		"/login",
		api.Login,
	},

	Route{
		"Register",
		strings.ToUpper("Post"),
		"/register",
		api.Register,
	},

	Route{
		"UpdateEvent",
		strings.ToUpper("Put"),
		"/update/{id}",
		api.Update,
	},

	Route{
		"Whoami",
		strings.ToUpper("Get"),
		"/whoami",
		api.Whoami,
	},
}
