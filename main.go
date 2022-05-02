package main

import (
	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
  
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
    	{ID: "3", Name: "Algorand Coin", Coin_Type: "ALG", Price: "4.5"},
	}
 
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world renter!",
		})
	})


	r.GET("/api/blockchain", func(c *gin.Context) {
		 c.IndentedJSON(200, prices)
	})
	
	r.GET("/api/blockchain/:id", func(c *gin.Context) {
			 
		    id := c.Param("id") 

		    for _, a := range prices {
		        if a.ID == id {
		            c.IndentedJSON(200, a)
		            return
		        }
		    }
		    c.IndentedJSON(500, gin.H{"message": "Server error id not found"})
	 
	})

	r.POST("/api/blockchain/price/create", func(c *gin.Context) {
			
			var newItem bcap_price
			if err := c.BindJSON(&newItem); err != nil {
		       c.IndentedJSON(500, gin.H{"message": "Server error item not created "})
	 
		    }
 
		    prices = append(prices, newItem)
		    c.IndentedJSON(500, newItem)
	
	})

	r.GET("/api/blockchain/price/update/", func(c *gin.Context) {
			
			var theItem bcap_price
			id := c.Param("id") 

			theItem = prices[id] 

			theItem.ID = c.Param("id")  
			theItem.Name = c.Param("name")  
			theItem.Coin_Type = c.Param("coin")  
			theItem.Price = c.Param("price")  
  			
  			prices[id] = theItem 
  			// itemsaved

			if err := c.BindJSON(&newItem); err != nil {
		       c.IndentedJSON(500, gin.H{"message": "Server error item not created "})	 
		    }
       
		    prices = append(prices, newItem)
		    c.IndentedJSON(500, newItem)
	})

	r.Run()
}


