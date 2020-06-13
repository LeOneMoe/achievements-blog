package factories

import (
	"github.com/LeOneMoe/go-gin-react-crud/dao/interfaces"
	"github.com/LeOneMoe/go-gin-react-crud/dao/mysql"
	"log"
)

func FactoryDAO(driver string) interfaces.UserDAO {
	var i interfaces.UserDAO

	switch driver {
	case "mysql":
		i = mysql.UserImplMysql{}
	default:
		log.Fatal(driver)
	}

	return i
}
