package parser

import (
	"github.com/Umiiii/solparser/types"
	"github.com/Umiiii/solparser/types/accounts"
	"github.com/gagliardetto/solana-go/rpc"
	"reflect"
)

func (s *SolParser) ProcessMeteoraSwapEvent(ix *rpc.ParsedInstruction) (*types.SwapTransactionEvent, error) {
	var acc accounts.MeteoraDLMMSwapAccounts
	var totalNumberOfCorrectAccount = reflect.TypeOf(acc).NumField()
	swapEvent := &types.SwapTransactionEvent{}
	if len(ix.Accounts) >= totalNumberOfCorrectAccount {
		acc = accounts.ParseAccountsIntoStruct[accounts.MeteoraDLMMSwapAccounts](ix.Accounts)
		swapEvent.PoolAddress = acc.LbPair.String()
	} else {
		return nil, nil
	}
	return swapEvent, nil
}
