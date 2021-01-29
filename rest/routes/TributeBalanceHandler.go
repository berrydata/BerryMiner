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
type TributeBalanceHandler struct {
}

//Incoming implementation for  handler
func (h *TributeBalanceHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.TributeBalanceKey)
	if err != nil {
		log.Printf("Problem reading TributeBalance from DB: %v\n", err)
		return 500, `{"error": "Could not read TributeBalance from DB"}`
	}
	return 200, fmt.Sprintf(`{ "TributeBalance": "%s" }`, string(v))
}
