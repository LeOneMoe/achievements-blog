package mysql

import "github.com/LeOneMoe/go-gin-react-crud/models"

type AchievementImplMysql struct {
}

func (dao AchievementImplMysql) Create(a *models.Achievement) error {
	query := "INSERT INTO achievements_blog.achievement (title, achievementText, date, authorID) VALUES(?, ?, ?, ?)"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(a.Title, a.AchievementText, a.UTCDate, a.AuthorID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	a.ID = id
	return nil
}

func (dao AchievementImplMysql) Update(id int64, a *models.Achievement) error {
	query := "UPDATE achievements_blog.achievement SET title = ?, achievementText = ?, date = ?, authorID = ? WHERE id = ?"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.AchievementText, a.UTCDate, a.AuthorID, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao AchievementImplMysql) Delete(id int64) error {
	query := "DELETE FROM achievements_blog.achievement WHERE id = ?"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (dao AchievementImplMysql) GetAll() ([]models.Achievement, error) {
	query := "SELECT id, title, achievementText, date, authorID FROM achievements_blog.achievement"
	var achievements []models.Achievement
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return achievements, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return achievements, err
	}

	for rows.Next() {
		var row models.Achievement
		err := rows.Scan(
			&row.ID, &row.Title,
			&row.AchievementText,
			&row.UTCDate,
			&row.AuthorID,
		)
		if err != nil {
			return nil, err
		}

		achievements = append(achievements, row)
	}

	return achievements, nil
}

func (dao AchievementImplMysql) GetByID(id int64) (models.Achievement, error) {
	query := "SELECT id, title, achievementText, date, authorID FROM achievements_blog.achievement WHERE id = ?"
	var achievement models.Achievement
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return achievement, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return achievement, err
	}

	row.Next()
	err = row.Scan(
		&achievement.ID,
		&achievement.Title,
		&achievement.AchievementText,
		&achievement.UTCDate,
		&achievement.AuthorID,
	)
	if err != nil {
		return achievement, err
	}

	return achievement, nil
}
