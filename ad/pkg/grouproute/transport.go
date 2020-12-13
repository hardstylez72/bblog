package grouproute

type insertRequest []params

func insertRequestConvert(r insertRequest) []params {
	return r
}

type listRequest struct {
	GroupId       int  `json:"groupId" validate:"required"`
	BelongToGroup bool `json:"belongToGroup"`
}

type listResponse []Route

func newListResponse(routes []Route) listResponse {
	return routes
}

type deleteRequest struct {
	Params []params `json:"groupId" validate:"dive"`
}
