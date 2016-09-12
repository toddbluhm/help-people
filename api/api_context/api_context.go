package api_context

import "github.com/toddbluhm/help-people-api/context"

type APIContext struct {
	*context.Context
	UserID string
}
