package route

type insertRequest struct {
	Route       string   `json:"route" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	Tags        []string `json:"tags"`
}

type updateRequest struct {
	Route       string   `json:"route" validate:"required"`
	Id          int      `json:"id" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Method      string   `json:"method" validate:"required"`
	Tags        []string `json:"tags"`
}

func insertRequestConvert(r *insertRequest) *Route {
	return &Route{
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
	}
}

func updateRequestConvert(r *updateRequest) *Route {
	return &Route{
		Id:          r.Id,
		Route:       r.Route,
		Description: r.Description,
		Method:      r.Method,
	}
}

type insertResponse RouteWithTags
type updateResponse Route

func newInsertResponse(group *RouteWithTags) *insertResponse {
	return (*insertResponse)(group)
}

func newUpdateResponse(route *Route) *updateResponse {
	return (*updateResponse)(route)
}

type listResponse []RouteWithTags

func newListResponse(routes []RouteWithTags) listResponse {
	return routes
}

type deleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type getRequest struct {
	Id int `json:"id" validate:"required"`
}
