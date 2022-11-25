package main

import (
	"marketeer/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// System
	router.GET("/api/marketeerHealth", api.HealthCheck)

	// User
	router.POST("/api/registerUser", api.RegisterUser) //api/createUser
	// router.POST("/api/updateUser", api.UpdateUser)
	// router.DELETE("/api/deleteUser", api.DeleteUser)

	// // Lookup
	// router.GET("/api/lookup/:number", api.RetrieveUser)

	router.Run(":80")
}
