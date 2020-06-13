package main

import (
	"github.com/LeOneMoe/go-gin-react-crud/handlers"
	"github.com/LeOneMoe/go-gin-react-crud/utilities"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = db.GetConnection(); err != nil {
		println(err)
	}

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
