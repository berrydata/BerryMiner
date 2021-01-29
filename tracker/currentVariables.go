package tracker

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	berryCommon "github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/config"
	"github.com/berrydata/BerryMiner/contracts"
	"github.com/berrydata/BerryMiner/db"
	"github.com/berrydata/BerryMiner/util"
)

var currentVarsLog = util.NewLogger("tracker", "CurrentVarsTracker")

//CurrentVariablesTracker concrete tracker type
type CurrentVariablesTracker struct {
}

func (b *CurrentVariablesTracker) String() string {
	return "CurrentVariablesTracker"
}

//Exec implementation for tracker
func (b *CurrentVariablesTracker) Exec(ctx context.Context) error {
	//cast client using type assertion since context holds generic interface{}
	DB := ctx.Value(berryCommon.DBContextKey).(db.DB)
	//get the single config instance
	cfg := config.GetConfig()

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(berryCommon.MasterContractContextKey).(*contracts.BerryMaster)
	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err := instance.GetCurrentVariables(nil)
	if err != nil {
		fmt.Println("Current Variables Retrieval Error")
		return err
	}

	//if we've mined it, don't save it
	myStatus, err := instance.DidMine(nil, currentChallenge, fromAddress)
	if err != nil {
		fmt.Println("My Status Retrieval Error")
		return err
	}
	bitSetVar := []byte{0}
	if myStatus {
		bitSetVar = []byte{1}
	}
	currentVarsLog.Info("Retrieved variables. challengeHash: %x", currentChallenge)

	err = DB.Put(db.CurrentChallengeKey, currentChallenge[:])
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(requestID)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.DifficultyKey, []byte(hexutil.EncodeBig(difficulty)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.QueryStringKey, []byte(queryString))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.GranularityKey, []byte(hexutil.EncodeBig(granularity)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}
	err = DB.Put(db.TotalTipKey, []byte(hexutil.EncodeBig(totalTip)))
	//if err != nil {
	//	fmt.Println("Current Variables Put Error")
	//	return err
	//}

	return DB.Put(db.MiningStatusKey, bitSetVar)
}
