package article

type DeleteArticleRequest struct {
	Article
}

type DeleteArticleResponse struct {
	Id string `json:"id"`
}

func NewDeleteArticleResponse(id string) *DeleteArticleResponse {
	return &DeleteArticleResponse{Id: id}
}
