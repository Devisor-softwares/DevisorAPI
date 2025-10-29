package handlers

import "github.com/gin-gonic/gin"

func RegisterHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "register"})
}

func LoginHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "login"})
}

func ListServersHandler(c *gin.Context) {
    c.JSON(200, gin.H{"servers": []string{}})
}

func ServerHandler(c *gin.Context) {
    c.String(200, "Server is running.")
}

// Map names to functions for dynamic registration
var HandlerMap = map[string]gin.HandlerFunc{
    "RegisterHandler": RegisterHandler,
    "LoginHandler": LoginHandler,
    "ListServersHandler": ListServersHandler,
	"ServerHandler": ServerHandler,
}
