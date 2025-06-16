package parser

import (
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/ink-web3/solparser-go/types"
	"github.com/ink-web3/solparser-go/types/accounts"
	"reflect"
)

func (s *SolParser) ParsePhoenixSwapEvent(ix *rpc.ParsedInstruction) (*types.SwapTransactionEvent, error) {
	var acc accounts.PhoenixSwapAccounts

	if len(ix.Accounts) != reflect.TypeOf(acc).NumField() {
		return nil, nil
	}
	acc = accounts.ParseAccountsIntoStruct[accounts.PhoenixSwapAccounts](ix.Accounts)
	swapEvent := &types.SwapTransactionEvent{}
	swapEvent.PoolAddress = acc.Market.String()

	return swapEvent, nil
}
