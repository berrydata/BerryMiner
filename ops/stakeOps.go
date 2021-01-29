package ops

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	berryCommon "github.com/berrydata/BerryMiner/common"
	berry "github.com/berrydata/BerryMiner/contracts"
	berry1 "github.com/berrydata/BerryMiner/contracts1"
	"github.com/berrydata/BerryMiner/util"
	"math/big"
	"time"
)

/**
 * This is the operational deposit component. Its purpose is to deposit Berry Tokens so you can mine
 */

func printStakeStatus(bigStatus *big.Int, started *big.Int) {
	//0-not Staked, 1=Staked, 2=LockedForWithdraw 3= OnDispute
	status := bigStatus.Uint64()
	stakeTime := time.Unix(started.Int64(), 0)
	switch status {
	case 0:
		fmt.Printf("Not currently staked\n")
	case 1:
		fmt.Printf("Staked in good standing since %s\n", stakeTime.UTC())
	case 2:
		startedRound := started.Int64()
		startedRound = ((startedRound + 86399) / 86400) * 86400
		target := time.Unix(startedRound, 0)
		timePassed := time.Now().Sub(target)
		delta := timePassed - (time.Hour * 24 * 7)
		if delta > 0 {
			fmt.Printf("Stake has been eligbile to withdraw for %s\n", delta)
		} else {
			fmt.Printf("Stake will be eligible to withdraw in %s\n", -delta)
		}
	case 3:
		fmt.Printf("Stake is currently under dispute")
	}
}

func Deposit(ctx context.Context) error {

	tmaster := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)

	publicAddress := ctx.Value(berryCommon.PublicAddress).(common.Address)
	balance, err := tmaster.BalanceOf(nil, publicAddress)
	if err != nil {
		return fmt.Errorf("couldn't get BRY balance: %s", err.Error())
	}

	status, startTime, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return fmt.Errorf("failed to get stake status: %s", err.Error())
	}

	if status.Uint64() != 0 && status.Uint64() != 2 {
		printStakeStatus(status, startTime)
		return nil
	}

	dat := crypto.Keccak256([]byte("stakeAmount"))
	var dat32 [32]byte
	copy(dat32[:], dat)
	stakeAmt, err := tmaster.GetUintVar(nil, dat32)
	if err != nil {
		return fmt.Errorf("fetching stake amount failed: %s", err.Error())
	}

	if balance.Cmp(stakeAmt) < 0 {
		return fmt.Errorf("insufficient balance (%s), mining stake requires %s BRY",
			util.FormatERC20Balance(balance),
			util.FormatERC20Balance(stakeAmt))
	}

	instance2 := ctx.Value(berryCommon.TransactorContractContextKey).(*berry1.BerryTransactor)
	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("couldn't prepare ethereum transaction: %s", err.Error())
	}

	tx, err := instance2.DepositStake(auth)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}
	fmt.Printf("Stake depositied with txn %s\n", tx.Hash().Hex())

	return nil
}

func ShowStatus(ctx context.Context) error {
	tmaster := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)

	publicAddress := ctx.Value(berryCommon.PublicAddress).(common.Address)
	status, startTime, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return fmt.Errorf("failed to get stake status: %s", err.Error())
	}

	printStakeStatus(status, startTime)
	return nil
}

func RequestStakingWithdraw(ctx context.Context) error {

	tmaster := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)
	publicAddress := ctx.Value(berryCommon.PublicAddress).(common.Address)
	status, startTime, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return fmt.Errorf("failed to get stake status: %s", err.Error())
	}
	if status.Uint64() != 1 {
		printStakeStatus(status, startTime)
		return nil
	}

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}

	instance2 := ctx.Value(berryCommon.TransactorContractContextKey).(*berry1.BerryTransactor)
	tx, err := instance2.RequestStakingWithdraw(auth)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}
	fmt.Printf("Withdrawl request sent with txn: %s\n", tx.Hash().Hex())

	return nil
}

func WithdrawStake(ctx context.Context) error {

	tmaster := ctx.Value(berryCommon.MasterContractContextKey).(*berry.BerryMaster)
	publicAddress := ctx.Value(berryCommon.PublicAddress).(common.Address)
	status, startTime, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return fmt.Errorf("failed to get stake status: %s", err.Error())
	}
	if status.Uint64() != 2 {
		fmt.Printf("Can't withdraw")
		printStakeStatus(status, startTime)
		return nil
	}

	auth, err := PrepareEthTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare ethereum transaction: %s", err.Error())
	}

	instance2 := ctx.Value(berryCommon.TransactorContractContextKey).(*berry1.BerryTransactor)
	tx, err := instance2.WithdrawStake(auth)
	if err != nil {
		return fmt.Errorf("contract failed: %s", err.Error())
	}
	fmt.Printf("Withdrew stake with txn %s\n", tx.Hash().Hex())
	return nil
}
