package router

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
)

// Route defines a valid endpoint with the type of action supported on it
type Route struct {
    Method string
    Path string
    Middleware func(http.Handler) http.Handler // Middleware function
    HandlerFunc http.HandlerFunc
}


// NewRouter returns a router handle loaded with all the supported routes
func NewRouter(routes[] Route) * chi.Mux {
    r: = chi.NewRouter()

    //  cors support
        r.Use(cors.Handler(cors.Options {
        AllowedOrigins: [] string {
            "*"
        },
        AllowedMethods: [] string {
            "GET", "POST", "PUT", "DELETE", "OPTIONS"
        },
        AllowedHeaders: [] string {
            "Accept", "Authorization", "Content-Type", "X-CSRF-Token"
        },
    }))

    for _,route: = range routes {
        handler: = route.HandlerFunc
        if route.Middleware != nil {
            handler = route.Middleware(handler)
        }
        r.Method(route.Method, route.Path, handler)
        log.Printf("Route added: %#v\n", route)
    }

        return r
}
