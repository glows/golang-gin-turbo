package request

import validation "github.com/go-ozzo/ozzo-validation"

type BasicOrganization struct {
	Title   string `json:"title" binding:"required"  example:"New Post"`
	Content string `json:"content" binding:"required"  example:"Lorem Ipsum"`
}

type CreateOrganizationRequest struct {
	*BasicPost
}

type UpdateOrganizationRequest struct {
	*BasicPost
}

func (bp BasicOrganization) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Title, validation.Required),
		validation.Field(&bp.Content, validation.Required),
	)
}
