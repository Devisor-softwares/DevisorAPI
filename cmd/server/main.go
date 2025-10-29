package main

import (
    "devisor/internal/config"
    "devisor/internal/db"
    // "devisor/internal/handlers"
    // "devisor/internal/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg, _ := config.LoadConfig()
    database, _ := db.ConnectDB(cfg.DBUrl)

    r := gin.Default()

    // r.POST("/register", handlers.Register(database))
    // r.POST("/login", handlers.Login(database, cfg.JWTSecret))

    // auth := r.Group("/api", middleware.AuthMiddleware(cfg.JWTSecret))
    // {
    //     auth.GET("/servers", handlers.ListServers(database))
    //     auth.POST("/servers", handlers.CreateServer(database))
    // }

    r.Run(":" + cfg.AppPort)
}
