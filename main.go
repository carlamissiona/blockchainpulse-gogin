package main

import (
	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
 
 //BLOCKPULSE_USER
 //BLOCKPULSE_PRICE
type bcap_user struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Email string  `json:"email"`    
}

type bcap_price struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Coin_Type string  `json:"coin_type"`
    Price string  `json:"price"`
}
var prices = []bcap_price{
    {ID: "1", Name: "Dogge Coin", Coin_Type: "DGC", Price: "56.99"},
    {ID: "2", Name: "Ethereum Coin", Coin_Type: "ETH", Price: "6.5"},
	}
 
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world renter!",
		})
	})


	r.GET("/api/blochain", func(c *gin.Context) {
		 c.IndentedJSON(200, prices)
	})
	r.GET("/api/blochain/:id", func(c *gin.Context) {
			 
		    id := c.Param("id") 

		    for _, a := range prices {
		        if a.ID == id {
		            c.IndentedJSON(200, a)
		            return
		        }
		    }
		    c.IndentedJSON(500, gin.H{"message": "Server error id not found"})
	 
	})


	r.Run()
}
