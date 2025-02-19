package response

import "basic_server/model"

type CreateOrganizationResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetOrganizationResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// type CollectionResponse struct {
// 	Collection interface{} `json:"collection"`
// 	Meta       Meta        `json:"meta"`
// }

// type Meta struct {
// 	Amount int `json:"amount"`
// }

func CreateOrganizationsCollectionResponse(posts []model.Post) CollectionResponse {
	collection := make([]GetPostResponse, 0)

	for index := range posts {
		collection = append(collection, GetPostResponse{
			ID:      posts[index].ID,
			Title:   posts[index].Title,
			Content: posts[index].Content,
		})
	}
	return CollectionResponse{Collection: collection, Meta: Meta{Amount: len(collection)}}
}
