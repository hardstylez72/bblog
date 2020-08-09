package article

type SaveArticleRequest struct {
	ArticleWithBody
}

type SaveArticleResponse struct {
	Id string `json:"id"`
}

func NewSaveArticleResponse(id string) *SaveArticleResponse {
	return &SaveArticleResponse{Id: id}
}
