package handler

import (
	"basic_server/model"
	"basic_server/request"
	"basic_server/response"
	"basic_server/service"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	OrganizationService service.OrganizationServiceI
}

func NewOrganizationHandler(OrganizationService service.OrganizationServiceI) *OrganizationHandler {
	return &OrganizationHandler{
		OrganizationService: OrganizationService,
	}
}

// GetOrganizationByID godoc
// @Summary Get organization by id
// @Description Get organization by id
// @ID get-organization
// @Tags Organizations Actions
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} response.GetOrganizationResponse
// @Failure 401 {object} response.Error
// @Security ApiKeyAuth
// @Router /organization/{id} [get]
func (handler OrganizationHandler) GetOrganizationByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	post := model.Post{}
	if err := handler.OrganizationService.GetByID(id, &post); err != nil {
		response.ErrorResponse(context, err.Status, "Server error")
		return
	}

	if post.ID == 0 {
		response.ErrorResponse(context, http.StatusNotFound, "Post not found")
		return
	}

	response.SuccessResponse(context, response.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}

// SaveOrganization godoc
// @Summary Create organization
// @Description Create post
// @ID posts-create
// @Tags Organizations Actions
// @Accept json
// @Produce json
// @Param params body request.CreateOrganizationRequest true "Post title and content"
// @Success 200 {string} response.CreatePostResponse
// @Failure 400 {string} string "Bad request"
// @Security ApiKeyAuth
// @Router /organizations [post]
func (handler OrganizationHandler) SaveOrganization(context *gin.Context) {
	var createPostRequest request.CreatePostRequest

	if err := context.ShouldBind(&createPostRequest); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	claims := jwt.ExtractClaims(context)
	id := claims["id"].(float64)

	newPost, restError := handler.OrganizationService.CreateOrganization(createPostRequest.Title, createPostRequest.Content, uint(id))
	if restError != nil {
		response.ErrorResponse(context, restError.Status, "Post can't be created")
		return
	}

	response.SuccessResponse(context, response.CreatePostResponse{
		ID:      newPost.ID,
		Title:   newPost.Title,
		Content: newPost.Content,
	})
}

// UpdateOrganization godoc
// @Summary Update organization
// @Description Update organization
// @ID organizations-update
// @Tags Organizations Actions
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Param params body request.UpdateOrganizationRequest true "Organization title and content"
// @Success 200 {string} response.GetOrganizationResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /organization/{id} [put]
func (handler OrganizationHandler) UpdateOrganization(context *gin.Context) {
	var updatePostRequest request.UpdatePostRequest

	if err := context.ShouldBind(&updatePostRequest); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	id, _ := strconv.Atoi(context.Param("id"))

	post := model.Post{}
	if err := handler.OrganizationService.GetByID(id, &post); err != nil {
		response.ErrorResponse(context, err.Status, "Server error")
		return
	}

	if post.ID == 0 {
		response.ErrorResponse(context, http.StatusNotFound, "Post not found")
		return
	}

	post.Title = updatePostRequest.Title
	post.Content = updatePostRequest.Content
	if err := handler.OrganizationService.Save(&post); err != nil {
		response.ErrorResponse(context, err.Status, "Data was not saved")
		return
	}

	response.SuccessResponse(context, response.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}

// GetOrganizations godoc
// @Summary Get all organizations
// @Description Get all organizations of all users
// @ID get-organizations
// @Tags Organizations Actions
// @Produce json
// @Success 200 {object} response.CollectionResponse
// @Failure 401 {object} response.Error
// @Security ApiKeyAuth
// @Router /organizations [get]
func (handler OrganizationHandler) GetOrganizations(context *gin.Context) {
	var posts []model.Post
	if err := handler.OrganizationService.GetAll(&posts); err != nil {
		response.ErrorResponse(context, http.StatusInternalServerError, "Server error")
		return
	}
	response.SuccessResponse(context, response.CreatePostsCollectionResponse(posts))
}

// DeleteOrganization godoc
// @Summary Delete organization
// @Description Delete organization
// @ID organizations-delete
// @Tags Organizations Actions
// @Param id path int true "Organization ID"
// @Success 200 {string} string "Organization deleted successfully"
// @Failure 404 {object} response.Error
// @Security ApiKeyAuth
// @Router /organization/{id} [delete]
func (handler OrganizationHandler) DeleteOrganization(context *gin.Context) {
	post := model.Post{}
	id, _ := strconv.Atoi(context.Param("id"))
	if err := handler.OrganizationService.GetByID(id, &post); err != nil {
		response.ErrorResponse(context, err.Status, "Server error")
		return
	}

	if post.ID == 0 {
		response.ErrorResponse(context, http.StatusNotFound, "Post not found")
		return
	}

	if err := handler.OrganizationService.Delete(&post); err != nil {
		response.ErrorResponse(context, err.Status, "Server error")
		return
	}

	response.SuccessResponse(context, "Post delete successfully")
}
