package lending

import "github.com/marcopollivier/techagenda/pkg/event"

type Props struct {
	Events []event.Event
}

type IndexRouteProps struct {
	InitialCount int `json:"initialCount"`
}
