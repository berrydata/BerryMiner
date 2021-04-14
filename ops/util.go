package ops

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	berryCommon "github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/rpc"
	"math/big"
)

func PrepareEthTransaction(ctx context.Context) (*bind.TransactOpts, error) {

	client := ctx.Value(berryCommon.ClientContextKey).(rpc.ETHClient)

	publicAddress := ctx.Value(berryCommon.PublicAddress).(common.Address)

	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return nil, fmt.Errorf("problem getting pending nonce: %+v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("problem getting gas price: %+v", err)
	}

	ethBalance, err := client.BalanceAt(context.Background(), publicAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("problem getting balance: %+v", err)
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, fmt.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	privateKey := ctx.Value(berryCommon.PrivateKey).(*ecdsa.PrivateKey)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
