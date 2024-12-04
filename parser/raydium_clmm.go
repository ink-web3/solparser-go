package parser

import (
	"fmt"

	"github.com/Umiiii/raydium-go/instructions"
	"github.com/Umiiii/raydium-go/types"
	types2 "github.com/Umiiii/solparser/types"
	"github.com/Umiiii/solparser/types/accounts"
	"github.com/gagliardetto/solana-go/rpc"
)

func (s *SolParser) ParseClmmSwapEvent(clmmSwapIx *rpc.ParsedInstruction) (*types2.SwapTransactionEvent, error) {
	swapEvent := &types2.SwapTransactionEvent{}
	acc := accounts.ParseAccountsIntoStruct[types.ClmmSwapSingleV2Accounts](clmmSwapIx.Accounts)
	_, err1 := instructions.DecodeRaydiumSwapV2Data(clmmSwapIx.Data)
	if err1 != nil {
		err := fmt.Errorf("error decoding clmm instruction: %v", err1)
		return nil, err
	}
	swapEvent.PoolAddress = acc.PoolState.String()

	return swapEvent, nil
}
