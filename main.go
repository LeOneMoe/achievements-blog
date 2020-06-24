package main

import (
	"github.com/LeOneMoe/go-gin-react-crud/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Don`t use this in prod!!!
	r.Use(cors.Default())

	prodAPI := r.Group("/api")
	{
		prodAPI.GET("/achievements", handlers.GetAllAchievements)
	}

	//devAPI := r.Group("/dev")
	//{
	//	devAPI.GET("/hi", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "hello world",
	//		})
	//	})
	//}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
