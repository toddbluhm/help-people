package api_context

import "github.com/toddbluhm/help-people-api/request_context"

type APIContext struct {
	*request_context.Context
	UserID string
}
