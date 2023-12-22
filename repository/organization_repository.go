package repository

import (
	"basic_server/model"

	"gorm.io/gorm"
)

type OrganizationRepositoryI interface {
	GetAll(posts *[]model.Post) error
	GetByID(id int, post *model.Post) error
	Create(post *model.Post) error
	Save(post *model.Post) error
	Delete(post *model.Post) error
}

type OrganizationRepository struct {
	DB *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{
		DB: db,
	}
}

func (repository *OrganizationRepository) GetAll(posts *[]model.Post) error {
	return repository.DB.Find(posts).Error
}

func (repository *OrganizationRepository) GetByID(id int, post *model.Post) error {
	return repository.DB.Where("id = ? ", id).Find(post).Error
}

func (repository *OrganizationRepository) Create(post *model.Post) error {
	return repository.DB.Create(post).Error
}

func (repository *OrganizationRepository) Save(post *model.Post) error {
	return repository.DB.Save(post).Error
}

func (repository *OrganizationRepository) Delete(post *model.Post) error {
	return repository.DB.Delete(post).Error
}
