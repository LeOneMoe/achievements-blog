package utilities

import (
	"encoding/json"
	"github.com/LeOneMoe/go-gin-react-crud/db"
	"os"
)

func GetConfiguration() (db.DataBase, error) {
	config := db.DataBase{}
	file, err := os.Open("./config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
