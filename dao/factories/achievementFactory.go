package factories

import (
	"github.com/LeOneMoe/go-gin-react-crud/dao/interfaces"
	"github.com/LeOneMoe/go-gin-react-crud/dao/mysql"
	"log"
)

func AchievementFactory(driver string) interfaces.AchievementDAO {
	var i interfaces.AchievementDAO

	switch driver {
	case "mysql":
		i = mysql.AchievementImplMysql{}
	default:
		log.Fatal(driver)
	}

	return i
}
