package needs

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api/api_context"
	"github.com/toddbluhm/help-people-api/models"
)

type NeedContext struct {
	*api_context.APIContext
}

func (ctx *NeedContext) GetNeeds(rw web.ResponseWriter, req *web.Request) {
	var needs []models.Need
	ctx.Context.DB.
		Select("*").
		From("needs").
		LoadStruct(&needs)
	data, err := json.Marshal(needs)
	if err != nil {
		panic("Could not marshal data")
	}
	fmt.Fprintf(rw, string(data))
}

func AttachRouter(rootRouter *web.Router) {
	rootRouter.
		Subrouter(NeedContext{}, "/needs").
		Get("/", (*NeedContext).GetNeeds)
}
