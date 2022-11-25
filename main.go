package main

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

func main() {
	ed25519Key := ed25519.GenPrivKey()
	fmt.Printf("Public Key: %v\n", base64.StdEncoding.EncodeToString(ed25519Key.PubKey().Bytes()))
	fmt.Printf("Private Key: %v\n", base64.StdEncoding.EncodeToString(ed25519Key.Key))

}
