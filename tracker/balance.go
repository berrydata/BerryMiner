package tracker

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	berryCommon "github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/config"
	"github.com/berrydata/BerryMiner/db"
	"github.com/berrydata/BerryMiner/rpc"
)

//BalanceTracker concrete tracker type
type BalanceTracker struct {
}

func (b *BalanceTracker) String() string {
	return "BalanceTracker"
}

//Exec implementation for tracker
func (b *BalanceTracker) Exec(ctx context.Context) error {

	//cast client using type assertion since context holds generic interface{}
	client := ctx.Value(berryCommon.ClientContextKey).(rpc.ETHClient)
	DB := ctx.Value(berryCommon.DBContextKey).(db.DB)

	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	balance, err := client.BalanceAt(ctx, fromAddress, nil)

	if err != nil {
		fmt.Println("balance Error, balance.go")
		return err
	}
	enc := hexutil.EncodeBig(balance)
	log.Printf("Balance: %v", enc)
	return DB.Put(db.BalanceKey, []byte(enc))
}
