package ops

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	berryCommon "github.com/berrydata/BerryMiner/common"
	berry "github.com/berrydata/BerryMiner/contracts"
	berry1 "github.com/berrydata/BerryMiner/contracts1"
	"github.com/berrydata/BerryMiner/rpc"
	"github.com/berrydata/BerryMiner/util"
	"log"
	"math/big"
)

/**
 * This is the operational transfer component. Its purpose is to transfer berry tokens
 */

func prepareTransfer(amt *big.Int,ctx context.Context) (*bind.TransactOpts, error) {
	instance := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)
	senderPubAddr := ctx.Value(berryCommon.PublicAddress).(common.Address)

	balance, err := instance.BalanceOf(nil, senderPubAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}
	fmt.Println("My balance", util.FormatERC20Balance(balance))
	if balance.Cmp(amt) < 0{
		return nil, fmt.Errorf("insufficent balance (%s BRY), requested %s BRY",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(amt))
	}
	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}
	return auth, nil
}

func Transfer(toAddress common.Address, amt *big.Int,ctx context.Context) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(berryCommon.TransactorContractContextKey).(*berry1.BerryTransactor)
	tx, err := instance2.Transfer(auth, toAddress, amt)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}
	fmt.Printf("Transferred %s to %s... with tx:\n%s\n", util.FormatERC20Balance(amt), toAddress.String()[:12], tx.Hash().Hex())
	return nil
}

func Approve(_spender common.Address, amt *big.Int,ctx context.Context) error {
	auth, err := prepareTransfer(amt, ctx)
	if err != nil {
		return err
	}

	instance2 := ctx.Value(berryCommon.TransactorContractContextKey).(*berry1.BerryTransactor)
	tx, err := instance2.Approve(auth,_spender,amt)
	if err != nil {
		return err
	}
	fmt.Printf("Approved %s to %s... with tx:\n%s\n", util.FormatERC20Balance(amt), _spender.String()[:12], tx.Hash().Hex())
	return nil
}


func Balance(ctx context.Context, addr common.Address) error {
	client := ctx.Value(berryCommon.ClientContextKey).(rpc.ETHClient)

	ethBalance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return fmt.Errorf("problem getting balance: %+v", err)
	}

	instance := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)
	trbBalance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%s\n", addr.String())
	fmt.Printf("%10s ETH\n",util.FormatERC20Balance(ethBalance))
	fmt.Printf("%10s BRY\n",util.FormatERC20Balance(trbBalance))
	return nil
}