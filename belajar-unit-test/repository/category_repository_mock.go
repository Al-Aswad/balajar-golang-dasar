package repository

import (
	"al-aswad/unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		return argument.Get(0).(*entity.Category)
	}

}
