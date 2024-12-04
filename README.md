# solparser
Solparser is a Go library designed to parse Solana swap transactions. 
It supports aggregator's swap transaction parsing.
Plus, a Solana RPC client is required to use this library.

# Usage
```go
package main

import (
    "context"
    "fmt"
    "github.com/Umiiii/solparser"
    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
)

func main() {
    // Initialize RPC client
    client := rpc.New("https://api.mainnet-beta.solana.com")
    uint64One := uint64(1)

    // Create parser instance
    p := parser.NewSolParser(client)
    
    // Transaction signature to parse
    sig := solana.MustSignatureFromBase58("5zrZnZa1bNawuJofcdPUu7ZnHF13xTuyeixoVS8Ev8MmfVZtZ5kNmxaSaiB9URxp57WAwzSV9zuma9KD5eHcxyvU")
    
    // Get parsed transaction
    opts := &rpc.GetParsedTransactionOpts{
        MaxSupportedTransactionVersion: &uint64One,
        Commitment: rpc.CommitmentConfirmed,
    }
    
    parsedTx, err := client.GetParsedTransaction(context.Background(), sig, opts)
    if err != nil {
        panic(err)
    }
    
    // Parse swap events
    events, err := p.ParseSwapEvent(parsedTx)
    if err != nil {
        panic(err)
    }
    
    // Process swap events
    for _, event := range events {
        fmt.Printf("Swap Event:\n")
        fmt.Printf("  Pool: %s\n", event.PoolAddress)
        fmt.Printf("  Market: %s\n", consts.ProgramToString(event.MarketProgramId))
        fmt.Printf("  Input Token: %s Amount: %s\n", event.InToken.Code, event.InToken.Amount)
        fmt.Printf("  Output Token: %s Amount: %s\n", event.OutToken.Code, event.OutToken.Amount)
    }
}
```

Example output:
```go
Swap Event 0 41:
    Pool: 4sB6mY3veRfRtszdVL4Lg8GHuz4NKwxZvxvhr3P22xBb
    Market: 675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8
    Input Token: G7BJzg55Afx6Tn6J9CJfA444jMf6QJbVjgQAevNTpump Amount: 10000000000
    Output Token: So11111111111111111111111111111111111111112 Amount: 160644272
Swap Event 1 49:
    Pool: 4wTV1YmiEkRvAtNtsSGPtUrqRYQMe5SKy2uB4Jjaxnjf
    Market: 6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P
    Input Token: So11111111111111111111111111111111111111112 Amount: 159037830
    Output Token: Go1d8YwLymjETPqxj4ybDg6JuAMCX1cQQULErYtgpump Amount: 2755053211659
```

# Supported DEXs
## Raydium
- V4 AMM (675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8)
- CLMM (CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK)
- CPMM (CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C)

## Orca
- Whirlpool (whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc)
- Token Swap V1 (DjVE6JNiYqPL2QXyCUUh8rNjHrbz9hXHNYt99MQ59qw1)
- Token Swap V2 (9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP)

## Other DEXes
- Pump.fun (6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P)
- Meteora DLMM (LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo)
- Phoenix (PhoeNiXZ8ByJGLkxNfZRnkUfjvmuYqLR89jjFHGqdXY)
- Lifinity V2 (2wT8Yq49kHgDzXuPxZSaeLaH1qbmGXtEyPy64bL7aD3c)


## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or support, please open an issue on the GitHub repository.