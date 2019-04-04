package services

import (
	"github.com/bennyzanuar/goblog/helpers"
	"github.com/bennyzanuar/goblog/models"
	"github.com/bennyzanuar/goblog/utils"
)

// Output struct
type Output struct {
	Result interface{}
	Error  interface{}
}

// FindAllPost func
func FindAllPost(page int, limit int) Output {
	// var posts models.Post
	posts := make([]*models.Post, 0)
	p := &helpers.ParamPagination{
		DB:      utils.GetDB(),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id"},
	}
	res, err := helpers.Pagging(p, &posts)
	return Output{Result: res, Error: err}
}
