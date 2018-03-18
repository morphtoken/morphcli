# Morphtoken CLI

## Examples

View current rates:

```
$ morphcli rates

|      |    BCH     |    BTC     |    DASH     |     ETH     |     LTC     |     XMR     |
+------+------------+------------+-------------+-------------+-------------+-------------+
| BCH  |          1 | 0.11235795 |  2.35885419 |  1.69615598 |  5.92205119 |  4.45181727 |
| BTC  | 8.76976066 |          1 | 20.99410136 | 15.09600328 | 52.70700637 | 39.62173809 |
| DASH | 0.41169701 | 0.04694507 |           1 |  0.70868290 |  2.47433400 |  1.86004519 |
| ETH  | 0.57221844 | 0.06524904 |  1.36984490 |           1 |  3.43908141 |  2.58528025 |
| LTC  | 0.16383061 | 0.01868131 |  0.39219729 |  0.28201310 |           1 |  0.74018593 |
| XMR  | 0.21784864 | 0.02484089 |  0.52151212 |  0.37499813 |  1.30928884 |           1 |
```

Start BTC -> XMR trade

```
$ morphcli exchange --input btc --refund yourbtcaddresshere --output xmr --address yourxmraddresshere

...

Visit https://morphtoken.com/morph/view?q=morphidhere for an easier view

------------ PENDING ------------
Waiting for a deposit, send BTC to depositaddresshere

Limits:
  Minimum amount accepted: 0.00300000 BTC
  Maximum amount accepted: 1.00000000 BTC

Send a single deposit. If the amount is outside the limits, a refund will happen.

==============================
Thank you for using MorphToken
```

Lookup an existing trade

```
$ morphcli view --id morphidhere
```


## Usage

```
$ morphcli --help

NAME:
   MorphToken CLI - Exchange coins instantly from your terminal

USAGE:
   morphcli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   morphtoken <contact@morphtoken.com>

COMMANDS:
     rates     Get all instant rates
     exchange  Exchange one coin for another
     view      Fetch an existing trade
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/morphtoken/morphcli
```

## Contribution

1. Fork ([https://github.com/morphtoken/morphcli/fork](https://github.com/morphtoken/morphcli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s -w .`
1. Create a new Pull Request

## Author

[morphtoken](https://github.com/morphtoken)
