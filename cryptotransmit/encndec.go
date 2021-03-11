package cryptotransmit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	rsmsc "github.com/rk0cc/restsnap-dl-lib/misc"
)

//Get public key
func getPubKey() string {
	ks, err := ioutil.ReadFile(rsmsc.GetDataPath("restsnap.key"))
	if err != nil {
		panic("Unable to read key from restsnap")
	}
	return string(ks)
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
