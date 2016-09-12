package need_types

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api/api_context"
)

var needTypes = []string{
	"Diapers",
	"Food",
	"Place to Sleep",
	"Flat Tire",
	"Someone to Listen",
}

type NeedTypesCtx struct {
	*api_context.APIContext
}

func (ctx *NeedTypesCtx) GetNeedTypes(rw web.ResponseWriter, req *web.Request) {
	data, err := json.Marshal(needTypes)
	if err != nil {
		panic("Could not marshal data")
	}
	fmt.Fprintf(rw, string(data))
}

func AttachRouter(rootRouter *web.Router) {
	rootRouter.Subrouter(NeedTypesCtx{}, "/").
		Get("/need-types", (*NeedTypesCtx).GetNeedTypes)
	return
}
