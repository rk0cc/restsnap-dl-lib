package cryptotransmit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"

	rsmsc "github.com/rk0cc/restsnap-dl-lib/misc"
)

//RSKeySet is a type allows JSON formatted to store unique data for crypto
//
//All data must be generated from Java
type RSKeySet struct {
	PrivateKey string `json:"private_key"` //Private key
	Append     string `json:"append"`      //String append
}

func getRSKey() RSKeySet {
	rsks := &RSKeySet{}
	rsksStr, err := ioutil.ReadFile(rsmsc.GetDataPath("restsnap.key"))
	if err != nil {
		panic("Unable to read key set from restsnap")
	}
	err = json.Unmarshal(rsksStr, rsks)
	if err != nil {
		panic("This may not RSKeySet format, may be modified manually")
	}
	return *rsks
}

//Get private key
func getPriKey() string {
	return string(getRSKey().PrivateKey)
}

func priKeyParser(prikStr string) *rsa.PrivateKey {
	blk, _ := pem.Decode([]byte(prikStr))
	if blk == nil {
		panic("Failed to parse block of private key")
	}
	prik, err := x509.ParsePKCS1PrivateKey(blk.Bytes)
	if err != nil {
		panic(err)
	}
	return prik
}

func pubKeyParser(pubkStr string) *rsa.PublicKey {
	blk, _ := pem.Decode([]byte(pubkStr))
	if blk == nil {
		panic("Failed to parse block of public key")
	}
	pubk, err := x509.ParsePKCS1PublicKey(blk.Bytes)
	if err != nil {
		panic(err)
	}
	return pubk
}
