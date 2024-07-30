package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madhu1997/lumel_assessment.git/config"
	"github.com/madhu1997/lumel_assessment.git/controllers"
	"github.com/madhu1997/lumel_assessment.git/db/initializers"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	go load_csv_data() //triggered data loading background
	r.GET("/api/v1/revenue_calculations", controllers.RevenueCalculation)
	r.POST("/api/v1/refresh_data", func(c *gin.Context) {
		go load_csv_data() // Trigger data refresh on demand
		c.JSON(http.StatusAccepted, gin.H{"message": "Data refresh triggered"})
	})

	r.Run()
}
