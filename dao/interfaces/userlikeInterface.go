package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/model"

type UserDAO interface {
	Create(m model.User)
	Read(m model.User)
	Update(m model.User)
	Delete(m model.User)
}
