package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "30770201010420a9923fa71e56280cc711bcfbd6d037509a2f979be4f17a6a54736141390fa745a00a06082a8648ce3d030107a1440342000497c48e38395e993f8e8c273c903104dcd6d3366e0ae6e2057417055f4e86f8b9a371dd691e8d98e6ff758ae16d8e1403f7150446c81eb0eb1f5e7d468e6c2d0e"
	testPayload string = "00094c1e71c89e8819a85f0f7e69968a35f53b2689644208b1ea1d05ae44f952"
	testSign    string = "9c4c94fe9839dbc86f1e5c386c26520b48fe88e909f70e5597280f61d8443fbee776db05c216e787973468dc1ca5715e86029ce5204a40650d38d59515fe6f6d"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromKey(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	w := makeTestWallet()
	tests := []test{
		{input: testPayload, ok: true},
		{input: testPayload, ok: false},
		{input: "10094c1e71c89e8819a85f0f7e69968a35f53b2689644208b1ea1d05ae44f952", ok: false},
	}
	for _, tc := range tests {
		ok := Verify(testSign, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("RestoreBigInts should return error when payload is not hex.")
	}
}
