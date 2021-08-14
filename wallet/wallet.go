package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/coseo12/nomacoin/utils"
)

const (
	signature     string = "9d325800c790b68a6ce28d65987df3403cf3c5a5228bb4d9f5dc0f8a1e078205249a4e8c616680ee23253b2b01a1431b94fbfe1764f61b1ad177b0ecdd9a9852"
	privateKey    string = "3077020101042034ce186c4d93c3e03e843911dabeddb0c408dae484adaef30222bbb3dc7e6eb8a00a06082a8648ce3d030107a144034200044cd3ad7260d1b4b1a7c2c93cc8c6bc6aeaad8f3e775e919eecc3ec404eb349e1bc3e57bc92a76b82f6a810e4d60ba08d8698eea2a1287483b610a8da356f8392"
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
)

func Start() {
	privateByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoredKey, err := x509.ParseECPrivateKey(privateByte)
	utils.HandleErr(err)

	signatureByte, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := signatureByte[:len(signatureByte)/2]
	sBytes := signatureByte[len(signatureByte)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
