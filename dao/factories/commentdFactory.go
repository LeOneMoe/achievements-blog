package factories

import (
	"github.com/LeOneMoe/go-gin-react-crud/dao/interfaces"
	"github.com/LeOneMoe/go-gin-react-crud/dao/mysql"
	"log"
)

func CommentFactory(driver string) interfaces.CommentDAO {
	var i interfaces.CommentDAO

	switch driver {
	case "mysql":
		i = mysql.CommentImplMysql{}
	default:
		log.Fatal(driver)
	}

	return i
}
