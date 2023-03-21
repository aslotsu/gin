package main

import (
	"github.com/gin-gonic/gin"
	"gonic-models/database"
	"gonic-models/routes"
	"log"
)

func init() {
	client := database.Connect()
	if client == nil {
		log.Println("There is an error with your mongo client")
	}
	if client != nil {

		log.Println("Your mongo client is working as expected")
	}

}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "You were successful in getting the Api 🇬🇭"})

	})

	router.POST("/new", routes.AddCountry())
	router.GET("/all", routes.GetCountries())
	router.GET("/all/sorted", routes.GetCountriesSorted())
	router.GET("/all/:id", routes.GetCountry())
	router.DELETE("/del/:id", routes.DeleteCountry())
	router.PUT("/repl/:id", routes.ReplaceCountry())
	router.PATCH("/upd/:id", routes.UpdateCountry())

	if err := router.Run(":8083"); err != nil {
		log.Println("Could not start router at said port", err)
		log.Fatal()
	}
}
