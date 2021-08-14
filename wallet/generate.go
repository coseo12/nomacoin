package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

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

func Verify() {

	privateByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privateByte)
	utils.HandleErr(err)

	signatureByte, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := signatureByte[:len(signatureByte)/2]
	sBytes := signatureByte[len(signatureByte)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)

	fmt.Println(ok)
}
