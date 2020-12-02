package group

type insertRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func insertRequestConvert(r *insertRequest) *Group {
	return &Group{
		Code:        r.Code,
		Description: r.Description,
	}
}

type insertResponse Group

func newInsertResponse(group *Group) *insertResponse {
	return (*insertResponse)(group)
}

type listResponse []Group

func newListResponse(groups []Group) listResponse {
	return groups
}

type deleteRequest struct {
	Id int `json:"id"`
}
