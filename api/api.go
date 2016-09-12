package api

import (
	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api/api_context"
	"github.com/toddbluhm/help-people-api/api/need_types"
)

func AttachRouter(rootRouter *web.Router) {
	router := rootRouter.Subrouter(api_context.APIContext{}, "/v1")
	need_types.AttachRouter(router)
	return
}
