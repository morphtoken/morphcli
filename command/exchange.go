package command

import (
    "log"
    "bytes"
    "strings"
    "net/http"
    "encoding/json"
    "github.com/urfave/cli"
)

type TradeRequest struct {
    Input TradeRequestInput     `json:"input"`
    Output []TradeRequestOutput `json:"output"`
}

type TradeRequestInput struct {
    Asset string                `json:"asset"`
    RefundAddress string        `json:"refund"`
}

type TradeRequestOutput struct {
    Asset string                `json:"asset"`
    Weight int16                `json:"weight"`
    Address string              `json:"address"`
}

type ErrorRequest struct {
    Code int16                  `json:"code"`
    Description string          `json:"description"`
    Success bool                `json:"success"`
}

func CmdExchange(c *cli.Context) error {
    inputAsset := strings.ToUpper(c.String("input"))
    outputAsset := strings.ToUpper(c.String("output"))
    outAddress := c.String("address")
    refundAddress := c.String("refund")
    if inputAsset == "" || outputAsset == "" || outAddress == "" || refundAddress == "" {
        log.Fatal("All arguments are required, check `exchange --help`")
    }

    reqInput := TradeRequestInput{Asset: inputAsset, RefundAddress: refundAddress}
    reqOutput := TradeRequestOutput{Asset: outputAsset, Address: outAddress, Weight: 10000}
    req := TradeRequest{Input: reqInput, Output: []TradeRequestOutput{reqOutput}}
    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(req)

    resp, err := http.Post("https://api.morphtoken.com/morph", "application/json", buf)
    if err != nil {
        log.Fatal("Failed to start trade: ", err)
        return nil
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Println("Failed to start trade")
        var errReq ErrorRequest
        if err := json.NewDecoder(resp.Body).Decode(&errReq); err == nil {
            log.Println("Reason:", errReq.Description, " [", errReq.Code, "]")
        }
        return nil
    }

    var trade Trade
    if err := json.NewDecoder(resp.Body).Decode(&trade); err != nil {
        log.Println(err)
        return nil
    }

    DisplayTrade(&trade)
	return nil
}
