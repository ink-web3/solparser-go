package parser

import (
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/ink-web3/solparser-go/types"
	"github.com/ink-web3/solparser-go/types/accounts"
	"reflect"
)

func (s *SolParser) ParseLifinitySwapEvent(ix *rpc.ParsedInstruction) (*types.SwapTransactionEvent, error) {
	var acc accounts.LifinitySwapV2Accounts
	if len(ix.Accounts) != reflect.TypeOf(acc).NumField() {
		return nil, fmt.Errorf("invalid number of accounts")
	}
	acc = accounts.ParseAccountsIntoStruct[accounts.LifinitySwapV2Accounts](ix.Accounts)
	swapEvent := &types.SwapTransactionEvent{}
	swapEvent.PoolAddress = acc.Amm.String()
	return swapEvent, nil
}
