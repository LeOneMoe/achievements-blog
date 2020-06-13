package mysql

import "github.com/LeOneMoe/go-gin-react-crud/models"

type UserImplMysql struct {
}

func (dao UserImplMysql) Read(m *models.User) error {
	panic("implement me")
}

func (dao UserImplMysql) Update(m *models.User) error {
	panic("implement me")
}

func (dao UserImplMysql) Delete(m *models.User) error {
	panic("implement me")
}

func (dao UserImplMysql) Create(u *models.User) error {
	query := "INSERT INTO achievements_blog.user (nickName, firstName, lastName) VALUES(?, ?, ?)"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.NickName, u.FirstName, u.LastName)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = id
	return nil
}
