package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/coseo12/nomacoin/utils"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	message := "I love you"

	hashedMessage := utils.Hash(message)

	hashAsByte, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsByte)
	utils.HandleErr(err)

	fmt.Printf("R:%d\nS:%d\n", r, s)
}
