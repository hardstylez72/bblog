package article

type UpdateArticleRequest struct {
	Id      string `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Body    string `json:"body" validate:"required"`
	UserId  string `json:"userId" validate:"required"`
	Preface string `json:"preface" validate:"required"`
}

type UpdateArticleResponse struct {
	Id string `json:"id"`
}

func NewUpdateArticleResponse(id string) *UpdateArticleResponse {
	return &UpdateArticleResponse{Id: id}
}
