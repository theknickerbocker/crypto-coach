# Crypto Coach
A CLI to get recommended crypto currency investments (legally not financial advice 😁).

# Installation
To install and use this CLI:
1. Ensure that [Go is installed](https://go.dev/doc/install).
2. Check that the GOPATH is added to a shell's PATH.
3. Run `go install github.com/theknickerbocker/crypto-coach@latest`.

After these steps the CLI should be installed. Run the following to verify:
```bash
❯ crypto-coach
```

# Usage
## `crypto-coach invest [command options] USD`
Get the recommended BTC and ETH investment (70/30 split) for given USD amount.

Example:
```bash
❯ crypto-coach invest 100
{"BTC":"0.001853856027156","ETH":"0.014665874709432"}
```

