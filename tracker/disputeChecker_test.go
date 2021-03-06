package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	berryCommon "github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/config"
	"github.com/berrydata/BerryMiner/db"
	"github.com/berrydata/BerryMiner/rpc"
)

func TestDisputeCheckerInRange(t *testing.T) {
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{ lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	if err != nil {
		t.Fatal(err)
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), berryCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, berryCommon.DBContextKey, DB)
	BuildIndexTrackers()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2*time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, berryCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err = disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB.Close()
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	cfg := config.GetConfig()
	cfg.DisputeThreshold = 0.000000001
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{ lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	if err != nil {
		t.Fatal(err)
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), berryCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, berryCommon.DBContextKey, DB)
	BuildIndexTrackers()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2*time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, berryCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err = disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB.Close()
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*IndexTracker) {
	for _, psr  := range psrs {
		//fmt.Print("\nIndex: ", psrIdx, psrs[psrIdx])
		err := psr.Exec(ctx)
		if err != nil {
			t.Fatalf("failed to execute psr: %v", err)
		}
	}
}
