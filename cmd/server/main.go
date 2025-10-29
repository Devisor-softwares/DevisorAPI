package main

import (
    "devisor/internal/config"
    // "devisor/internal/db"
    // "devisor/internal/handlers"
    // "devisor/internal/middleware"
    "devisor/internal/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    routesCfg, err := config.LoadRoutes("routes.yml")
    if err != nil {
        panic(err)
    }

    for _, route := range routesCfg.Routes {
        h, ok := handlers.HandlerMap[route.Handler]
        if !ok {
            panic("handler not found: " + route.Handler)
        }

        switch route.Method {
        case "GET":
            if route.Auth {
                r.GET(route.Path, AuthMiddleware(), h)
            } else {
                r.GET(route.Path, h)
            }
        case "POST":
            if route.Auth {
                r.POST(route.Path, AuthMiddleware(), h)
            } else {
                r.POST(route.Path, h)
            }
        }
    }

    r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: check auth here
        c.Next()
    }
}
