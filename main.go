package main

import (
	"github.com/LeOneMoe/go-gin-react-crud/handlers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	//cfg, err := utilities.GetConfiguration()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//userDAO := factories.UserFactory(cfg.Driver)

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	standardAPI := r.Group("")
	{
		standardAPI.GET("/ping", handlers.PostPingPong)
	}

	devAPI := r.Group("/dev")
	{
		devAPI.GET("/hi", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
