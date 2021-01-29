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
type MiningStatusHandler struct {
}

//Incoming implementation for  handler
func (h *MiningStatusHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.MiningStatusKey)
	if err != nil {
		log.Printf("Problem reading MiningStatus from DB: %v\n", err)
		return 500, `{"error": "Could not read CurrentChallenge from DB"}`
	}
	return 200, fmt.Sprintf(`{ "miningStatus": "%s" }`, string(v))
}
