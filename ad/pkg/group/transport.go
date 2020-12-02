package group

type insertGroupRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func insertGroupRequestConvert(r *insertGroupRequest) *Group {
	return &Group{
		Code:        r.Code,
		Description: r.Description,
	}
}

type insertGroupResponse Group

func newInsertGroupResponse(group *Group) *insertGroupResponse {
	return (*insertGroupResponse)(group)
}

type getGroupsResponse []Group

func newGetGroupsResponse(groups []Group) getGroupsResponse {
	return groups
}
