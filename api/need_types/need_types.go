package need_types

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api/api_context"
)

func GetNeedTypes(ctx *api_context.APIContext, rw web.ResponseWriter, req *web.Request) {
	var needTypes []string
	ctx.Context.DB.
		Select("name").
		From("need_types").
		LoadValues(&needTypes)
	data, err := json.Marshal(needTypes)
	if err != nil {
		panic("Could not marshal data")
	}
	fmt.Fprintf(rw, string(data))
}

func AttachRouter(rootRouter *web.Router) {
	rootRouter.Get("/need-types", GetNeedTypes)
}
