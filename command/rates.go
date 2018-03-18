package command

import (
    "os"
    "log"
    "fmt"
    "sort"
    "net/http"
    "encoding/json"
    "github.com/urfave/cli"
    "github.com/olekukonko/tablewriter"
)

type Rates struct {
    Timestamp float32   `json:"timestamp"`
    Type string         `json:"type"`
    Data map[string]map[string]string  `json:"data"`
}

func CmdRates(c *cli.Context) error {
    resp, err := http.Get("https://api.morphtoken.com/rates")
    if err != nil {
        log.Fatal("Failed to get rates: ", err)
        return nil
    }

    defer resp.Body.Close()

    var rates Rates
    if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
        log.Println(err)
        return nil
    }

    var keys []string
    for k := range rates.Data {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    table := tablewriter.NewWriter(os.Stdout)
    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    var header []string = append([]string{""}, keys...)
    table.SetHeader(header)
    for _, k := range keys {
        var values[]string
        values = append(values, k)
        for _, other := range keys {
            var val = rates.Data[k][other]
            if val != "" {
              values = append(values, val)
            } else {
              values = append(values, "1")
            }
        }
        table.Append(values)
    }
    fmt.Println("")
    table.Render()
    fmt.Println("")

	return nil
}
