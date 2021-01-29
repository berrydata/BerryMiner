package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/db"
)

//BalanceHandler handles balance requests
type QueryStringHandler struct {
}

//Incoming implementation for  handler
func (h *QueryStringHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.QueryStringKey)
	if err != nil {
		log.Printf("Problem reading QueryString from DB: %v\n", err)
		return 500, `{"error": "Could not read QueryString from DB"}`
	}
	return 200, fmt.Sprintf(`{ "QueryString": "%s" }`, string(v))
}
