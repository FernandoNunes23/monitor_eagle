package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "api/server/router/memory"
    "api/server/router/monit"
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

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/memory",
        memory.GetMemoryInfo,
    },
    Route{
        "Index",
        "GET",
        "/swap",
        memory.GetSwapInfo,
    },
    Route{
        "Index",
        "GET",
        "/monit",
        monit.ShowCreatedMonits,
    },
    Route{
        "Index",
        "GET",
        "/monit/{name}",
        monit.GetMonitInfo,
    },
    Route{
        "Index",
        "GET",
        "/monit/{name}/launch",
        monit.LaunchMonit,
    },
    Route{
        "Index",
        "POST",
        "/monit/create",
        monit.CreateMonit,
    },
}