package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type CommentDAO interface {
	Create(m models.Comment)
	Read(m models.Comment)
	Update(m models.Comment)
	Delete(m models.Comment)
}
