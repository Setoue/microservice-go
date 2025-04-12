package services

import (
	"log"

	"github.com/Setoue/microservice-go-category/internal/entities"
)

type createCategoryServices struct {
	//db
}

func NewCreateCategoryServices() *createCategoryServices {
	return &createCategoryServices{
		//db: db,
	}
}

func (s *createCategoryServices) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: save category to db
	log.Println(category)

	return nil
}
