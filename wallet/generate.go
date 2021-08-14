package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/coseo12/nomacoin/utils"
)

func Generate() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.HandleErr(err)

	fmt.Printf("%x\n\n\n\n\n", keyAsBytes)

	message := "I love you"

	hashedMessage := utils.Hash(message)

	hashAsByte, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsByte)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("%x\n", signature)
}
