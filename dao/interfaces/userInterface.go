package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type UserDAO interface {
	Create(u *models.User) error
	Read(u *models.User) error
	Update(u *models.User) error
	Delete(u *models.User) error
}
