package service

import (
	"al-aswad/unit-test/entity"
	"al-aswad/unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryREpository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
