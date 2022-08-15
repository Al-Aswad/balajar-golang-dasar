package repository

import "al-aswad/unit-test/entity"

type CategoryREpository interface {
	FindById(id string) *entity.Category
}
