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
type GranularityHandler struct {
}

//Incoming implementation for  handler
func (h *GranularityHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.GranularityKey)
	if err != nil {
		log.Printf("Problem reading Granularity from DB: %v\n", err)
		return 500, `{"error": "Could not read Granularity from DB"}`
	}
	return 200, fmt.Sprintf(`{ "Granularity": "%s" }`, string(v))
}
