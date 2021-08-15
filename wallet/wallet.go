package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/coseo12/nomacoin/utils"
)

const (
	filename string = "nomadcoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func createPrivatekey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(filename, bytes, 0644)
	utils.HandleErr(err)
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			fmt.Println("??")
			// yes -> restore from file
		} else {
			key := createPrivatekey()
			persistKey(key)
			w.privateKey = key
		}
	}
	return w
}
