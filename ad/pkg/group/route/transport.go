package route

type insertRequest []insertParams

func insertRequestConvert(r insertRequest) []insertParams {
	return r
}

type insertResponse []Route

type listRequest struct {
	GroupId int `json:"id"`
}

type listResponse []Route

func newListResponse(routes []Route) listResponse {
	return routes
}

type deleteRequest struct {
	Id int `json:"id"`
}
