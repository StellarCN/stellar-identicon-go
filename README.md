# Stellar Identicon Generator (Golang)

The go implementation of [SEP-0033](https://github.com/stellar/stellar-protocol/blob/master/ecosystem/sep-0033.md), 
which you can use to generate identicons for Stellar wallets - unique icons, generated based on the wallet public key.

## Install
```shell
go get -u github.com/overcat/stellar-identicon-go
```

## Usage
```go
import identicon "github.com/overcat/stellar-identicon-go"

func main() {
	publicKey := "GAQL4ZLRIJBGSCXE6XC4XPZ5W6FGCJHXLAMB4M7ZQ52HFPDTL6GSVP4W"
	img, err := identicon.Generate(publicKey, identicon.Width, identicon.Height)
	if err != nil {
		
	}
}
```

## Thanks
This project was inspired by [LOBSTR](https://github.com/Lobstrco).
