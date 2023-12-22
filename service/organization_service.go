package service

import (
	"basic_server/model"
	"basic_server/repository"
	"net/http"
)

type OrganizationServiceI interface {
	CreateOrganization(title, content string, userID uint) (*model.Post, *RestError)
	GetAll(posts *[]model.Post) *RestError
	GetByID(id int, post *model.Post) *RestError
	Create(post *model.Post) *RestError
	Save(post *model.Post) *RestError
	Delete(post *model.Post) *RestError
}

type OrganizationService struct {
	OrgRepository repository.OrganizationRepositoryI
}

func (service *OrganizationService) GetAll(posts *[]model.Post) *RestError {
	if err := service.OrgRepository.GetAll(posts); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

func (service *OrganizationService) GetByID(id int, post *model.Post) *RestError {
	if err := service.OrgRepository.GetByID(id, post); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

func (service *OrganizationService) Create(post *model.Post) *RestError {
	if err := service.OrgRepository.Create(post); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

func (service *OrganizationService) Save(post *model.Post) *RestError {
	if err := service.OrgRepository.Save(post); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

func (service *OrganizationService) Delete(post *model.Post) *RestError {
	if err := service.OrgRepository.Delete(post); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

func (service *OrganizationService) CreatePost(title, content string, userID uint) (*model.Post, *RestError) {
	post := &model.Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	if err := service.OrgRepository.Create(post); err != nil {
		return nil, &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	return post, nil
}

func NewOrganizationService(OrgRepo repository.OrganizationRepositoryI) PostServiceI {
	return &OrganizationService{OrgRepository: OrgRepo}
}
