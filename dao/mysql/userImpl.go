package mysql

import "github.com/LeOneMoe/go-gin-react-crud/models"

type UserImplMysql struct {
}

func (dao UserImplMysql) Create(user *models.User) error {
	query := "INSERT INTO achievements_blog.user (nickName, firstName, lastName) VALUES(?, ?, ?)"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.NickName, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func (dao UserImplMysql) Update(id int64, u *models.User) error {
	query := "UPDATE achievements_blog.user SET nickName = ?, firstName = ?, lastName = ? WHERE id = ?"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.NickName, u.FirstName, u.LastName, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao UserImplMysql) Delete(id int64) error {
	query := "DELETE FROM achievements_blog.user WHERE id = ?"
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

func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT id, nickName, firstName, lastName FROM achievements_blog.user"
	var users []models.User
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.NickName, &row.FirstName, &row.LastName)
		if err != nil {
			return nil, err
		}

		users = append(users, row)
	}

	return users, nil
}

func (dao UserImplMysql) GetByID(id int64) (models.User, error) {
	query := "SELECT id, nickName, firstName, lastName FROM achievements_blog.user WHERE id = ?"
	var user models.User
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return user, err
	}

	row.Next()
	err = row.Scan(&user.ID, &user.NickName, &user.FirstName, &user.LastName)
	if err != nil {
		return user, err
	}

	return user, nil
}
