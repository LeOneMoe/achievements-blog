package factories

import (
	"github.com/LeOneMoe/go-gin-react-crud/dao/interfaces"
	"github.com/LeOneMoe/go-gin-react-crud/dao/mysql"
	"log"
)

func LikeFactory(driver string) interfaces.LikeDAO {
	var i interfaces.LikeDAO

	switch driver {
	case "mysql":
		i = mysql.LikeImplMysql{}
	default:
		log.Fatal(driver)
	}

	return i
}
