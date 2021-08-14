package wallet

import (
	"crypto/ecdsa"
	"os"
)

const (
	signature     string = "a30f035c8d44dea5c1355a963e18e0a9426622cb3467f8b553538099ffab6fabc775b6d9a9014cd01be0436c980cf6deec7e4cacaa0258d7ecc3e77c329ab773"
	privateKey    string = "307702010104204eb9438b09d338be90fdc55d9391d4d5e10f620ca458818a3941e0bd400c2514a00a06082a8648ce3d030107a14403420004cf39c60bcf171dfdb10d9d51345558c522169e23f4c148c80658cac50303f0e7d410fe4990f2835793ab6fc5f53d81260346deb6dc38d2041f35d88c31f7fd9c"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("nomadcoin.wallet")
	return os.IsNotExist(err)
}

func Wallet() *wallet {
	if w == nil {
		if hasWalletFile() {
			// yes -> restore from file
		}
		// no -> crate prv key, save to file
	}
	return w
}
