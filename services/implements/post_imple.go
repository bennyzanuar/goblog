package implements

import (
	"github.com/bennyzanuar/goblog/models"
	"github.com/bennyzanuar/goblog/services/interfaces"
	"github.com/bennyzanuar/goblog/utils"
)

func FindAll() interfaces.Output {
	var posts models.Post
	err := utils.GetDB().Find(&posts).Error
	if err != nil {
		return interfaces.Output{Error: err}
	}
	return interfaces.Output{Result: posts}
}
