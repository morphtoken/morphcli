package command

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

func CmdView(c *cli.Context) error {
	tradeId := c.String("id")
	if tradeId == "" {
		log.Fatal("ID not specified")
		return nil
	}

	resp, err := http.Get(fmt.Sprintf("https://api.morphtoken.com/morph/%s", tradeId))
	if err != nil {
		log.Fatal("Failed to get trade: ", err)
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			log.Fatal("Trade not found")
		}
		log.Fatal("Failed to load trade")
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
