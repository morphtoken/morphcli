package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

type Trade struct {
	CreatedAt       string        `json:"created_at"`
	ID              string        `json:"id"`
	State           string        `json:"state"`
	Input           TradeInput    `json:"input"`
	Output          []TradeOutput `json:"output"`
	Refund          TradeRefund   `json:"refund"`
	RemainingWeight int16         `json:"remaining_weight"`
	Tag             string        `json:"tag"`
}

type TradeInput struct {
	Asset                string      `json:"asset"`
	Received             json.Number `json:"received"`
	ConfirmedHeight      int32       `json:"confirmed_at_height"`
	DepositAddress       string      `json:"deposit_address"`
	LegacyDepositAddress string      `json:"legacy_deposit_address"`
	RefundAddress        string      `json:"refund_address"`
	Limits               struct {
		Min json.Number `json:"min"`
		Max json.Number `json:"max"`
	} `json:"limits"`
}

type TradeOutput struct {
	Asset           string      `json:"asset"`
	Weight          int16       `json:"weight"`
	Address         string      `json:"address"`
	SeenRate        string      `json:"seen_rate"`
	FinalRate       string      `json:"final_rate"`
	ConvertedAmount json.Number `json:"converted_amount"`
	TxID            string      `json:"txid"`
	NetworkFee      struct {
		Flat bool        `json:"flat"`
		Fee  json.Number `json:"fee"`
	} `json:"network_fee"`
}

type TradeRefund struct {
	Asset       string      `json:"asset"`
	NetworkFee  json.Number `json:"network_fee"`
	FinalAmount json.Number `json:"final_amount"`
	TxID        string      `json:"txid"`
}

func FormatAmount(asset string, amount json.Number) string {
	num := new(big.Int)
	num, success := num.SetString(amount.String(), 10)
	if !success {
		log.Println("Failed to convert number")
	}
	float := new(big.Float).SetInt(num)
	factor := new(big.Int)
	switch asset {
	case "ETH":
		factor, _ = factor.SetString("1000000000000000000", 10)
	case "XMR":
		factor, _ = factor.SetString("1000000000000", 10)
	default:
		factor, _ = factor.SetString("100000000", 10)
	}
	div := new(big.Float).Quo(float, new(big.Float).SetInt(factor))
	return fmt.Sprintf("%.6f", div)
}

func DisplayTrade(trade *Trade) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "   ")
	err := encoder.Encode(trade)
	if err != nil {
		log.Fatal("Failed to parse trade in JSON", err)
		return
	}

	fmt.Println("Trade data:")
	fmt.Println(buffer.String())

	fmt.Printf("\nVisit https://morphtoken.com/morph/view?q=%s for an easier view\n\n", trade.ID)
	fmt.Printf("------------ %s ------------\n", trade.State)
	if trade.State == "PENDING" {
		fmt.Printf("Waiting for a deposit, send %s to %s\n\n", trade.Input.Asset, trade.Input.DepositAddress)
		fmt.Printf("Limits:\n")
		fmt.Printf("  Minimum amount accepted: %s %s\n",
			FormatAmount(trade.Input.Asset, trade.Input.Limits.Min), trade.Input.Asset)
		fmt.Printf("  Maximum amount accepted: %s %s\n\n",
			FormatAmount(trade.Input.Asset, trade.Input.Limits.Max), trade.Input.Asset)
		fmt.Println("Send a single deposit. If the amount is outside the limits, a refund will happen.")
	}
	fmt.Println("\n==============================")
	fmt.Printf("Thank you for using MorphToken\n\n")
}
