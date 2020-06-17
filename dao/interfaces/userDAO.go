package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type UserDAO interface {
	Create(u *models.User) error
	Update(id int64, u *models.User) error
	Delete(id int64) error
	GetAll() ([]models.User, error)
	GetByID(id int64) (models.User, error)
}
