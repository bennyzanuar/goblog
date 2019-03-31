package implements

import (
	"github.com/bennyzanuar/goblog/models"
	"github.com/bennyzanuar/goblog/utils"
)

// Save(*models.Post) Output
// Delete(*models.Post) Output
// FindByID(int) Output
// FindAll() Output

func FindAll() Output {
	// utils.GetDB().
	var posts models.Post
	// posts := make([]*models.Post, 0)
	err := utils.GetDB().Find(&posts).Error

	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: posts}
}
