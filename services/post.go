package services

import (
	"github.com/bennyzanuar/goblog/helpers"
	"github.com/bennyzanuar/goblog/models"
	"github.com/bennyzanuar/goblog/utils"
)

type ParamsOutput struct {
	Count    int `json:"count"`
	Pages    int `json:"pages"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PrevPage int `json:"prevPage"`
	NextPage int `json:"nextPage"`
}

// Output struct
type Output struct {
	ParamsOutput
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

	output := Output{
		ParamsOutput: ParamsOutput{
			Count:    res.Count,
			Pages:    res.Pages,
			Offset:   res.Offset,
			Limit:    res.Limit,
			Page:     res.Page,
			PrevPage: res.PrevPage,
			NextPage: res.NextPage,
		},
		Result: res.Records,
		Error:  err,
	}
	return output
}
