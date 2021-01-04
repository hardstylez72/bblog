package userroute

import (
	"github.com/hardstylez72/bblog/ad/pkg/group"
	"github.com/hardstylez72/bblog/ad/pkg/route"
)

type Route route.Route

type RouteWithGroups struct {
	Groups []group.Group `json:"groups"`
	Route
}
