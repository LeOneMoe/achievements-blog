package mysql

import "github.com/LeOneMoe/go-gin-react-crud/models"

type CommentImplMysql struct {
}

func (dao CommentImplMysql) Create(a *models.Comment) error {
	query := "INSERT INTO achievements_blog.comment (userID, achievementID, commentText) VALUES(?, ?, ?)"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		a.UserID,
		a.AchievementID,
		a.CommentText,
	)
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

func (dao CommentImplMysql) Update(id int64, a *models.Comment) error {
	query := "UPDATE achievements_blog.comment SET userID = ?, achievementID = ?, commentText = ? WHERE id = ?"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		a.UserID,
		a.AchievementID,
		a.CommentText,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (dao CommentImplMysql) Delete(id int64) error {
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

func (dao CommentImplMysql) GetAllByPost(achID int64) ([]models.Comment, error) {
	query := "SELECT id, userID, achievementID, commentText FROM achievements_blog.comment WHERE achievementID = ?"
	var comments []models.Comment
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return comments, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(achID)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var row models.Comment
		err := rows.Scan(
			&row.ID,
			&row.UserID,
			&row.AchievementID,
			&row.CommentText,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, row)
	}

	return comments, nil
}

func (dao CommentImplMysql) GetCount(achID int64) (int, error) {
	query := "SELECT COUNT(*) as count FROM achievements_blog.comment WHERE achievementID = ?"
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
