package mysql

import (
	"github.com/LeOneMoe/go-gin-react-crud/models"
)

type LikeImplMysql struct {
}

func (l2 LikeImplMysql) Create(l *models.Like) error {
	query := "INSERT INTO achievements_blog.`like` (userID, achievementID) VALUES(?, ?)"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		l.UserID,
		l.AchievementID,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	l.ID = id
	return nil
}

func (l2 LikeImplMysql) Delete(achID, userID int64) error {
	query := "DELETE FROM achievements_blog.`like` WHERE userID = ? and achievementID = ?"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(achID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (l2 LikeImplMysql) GetCount(achID int64) (int, error) {
	query := "SELECT COUNT(*) as count FROM achievements_blog.`like` WHERE achievementID = ?"
	var count int
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row, err := stmt.Query(achID)
	if err != nil {
		return -1, err
	}

	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}
