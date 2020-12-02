package route

type insertRequest struct {
	Route       string `json:"route"`
	Description string `json:"description"`
	Method      string `json:"method"`
}

func insertRequestConvert(r *insertRequest) *Route {
	return &Route{
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
	}
}

type insertResponse Route

func newInsertResponse(group *Route) *insertResponse {
	return (*insertResponse)(group)
}

type listResponse []Route

func newListResponse(groups []Route) listResponse {
	return groups
}

type deleteRequest struct {
	Id int `json:"id"`
}
