package cryptotransmit

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"

	"golang.org/x/crypto/sha3"
)

//GenDataSign : Generate data signature for varify
//Note: This will return JSON formatted byte array
func GenDataSign(data []byte) []byte {
	hashedData := sha3.New256()
	if _, err := hashedData.Write(data); err != nil {
		panic(err)
	}
	sum := hashedData.Sum([]byte(getRSKey().Append))
	sign, err := rsa.SignPKCS1v15(rand.Reader, priKeyParser(getPriKey()), crypto.SHA3_256, sum)
	if err != nil {
		panic("Create new signature failed")
	}
	marsedData, _ := json.Marshal(struct {
		Sign    string `json:"sign"`
		HashSum string `json:"hashsum"`
	}{
		Sign:    string(sign),
		HashSum: string(sum),
	})
	return marsedData
}

//VerifyData : Verify signature that is generated before
func VerifyData(hashSum, sign []byte, publicKey string) bool {
	failed := rsa.VerifyPKCS1v15(pubKeyParser(publicKey), crypto.SHA3_256, hashSum, sign)
	if failed != nil {
		return false
	}
	return true
}
