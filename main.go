package main

import (
	"github.com/LeOneMoe/go-gin-react-crud/dao/factories"
	"github.com/LeOneMoe/go-gin-react-crud/handlers"
	"github.com/LeOneMoe/go-gin-react-crud/models"
	"github.com/LeOneMoe/go-gin-react-crud/utilities"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	userDAO := factories.FactoryDAO(config.Driver)

	user := models.User{
		NickName:  "LeOneMoe",
		FirstName: "Oleg",
		LastName:  "Vzukov",
	}

	err = userDAO.Create(&user)

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	standardAPI := r.Group("")
	{
		standardAPI.GET("/ping", handlers.PostPingPong)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
