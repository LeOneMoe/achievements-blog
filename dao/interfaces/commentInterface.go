package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/model"

type CommentDAO interface {
	Create(m model.Comment)
	Read(m model.Comment)
	Update(m model.Comment)
	Delete(m model.Comment)
}
