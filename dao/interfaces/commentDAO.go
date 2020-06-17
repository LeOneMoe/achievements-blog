package interfaces

import "github.com/LeOneMoe/go-gin-react-crud/models"

type CommentDAO interface {
	Create(c *models.Comment) error
	Update(id int64, c *models.Comment) error
	Delete(id int64) error
	GetAllByPost(achID int64) ([]models.Comment, error)
	GetCount(achID int64) (int, error)
}
