package dataServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	berryCommon "github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/config"
	"github.com/berrydata/BerryMiner/contracts"
	"github.com/berrydata/BerryMiner/db"
	"github.com/berrydata/BerryMiner/rpc"
)

func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	cfg := config.GetConfig()
	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewBerryMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating berry master instance: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), berryCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, berryCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, berryCommon.MasterContractContextKey, masterInstance)
	ds, err := CreateServer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ds.Start(ctx, exitCh)

	time.Sleep(5000 * time.Millisecond)

	resp, err := http.Get("http://localhost:5000/balance")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Finished: %+v", resp)
	exitCh <- 1
	time.Sleep(1 * time.Second)
	if !ds.Stopped {
		t.Fatal("Did not stop server")
	}
}
