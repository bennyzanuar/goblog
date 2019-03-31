package interfaces

import "github.com/bennyzanuar/goblog/models"

// Output struct
type Output struct {
	Result interface{}
	Error  error
}

// PostInter interfaces
type PostInter interface {
	Save(*models.Post) Output
	Delete(*models.Post) Output
	FindByID(int) Output
	FindAll() Output
}
