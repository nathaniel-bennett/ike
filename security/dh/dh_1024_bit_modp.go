package dh

import (
	"math/big"

	"github.com/nathaniel-bennett/ike/message"
)

const (
	// Parameters
	Group2PrimeString string = "FFFFFFFFFFFFFFFFC90FDAA22168C234" +
		"C4C6628B80DC1CD129024E088A67CC74" +
		"020BBEA63B139B22514A08798E3404DD" +
		"EF9519B3CD3A431B302B0A6DF25F1437" +
		"4FE1356D6D51C245E485B576625E7EC6" +
		"F44C42E9A637ED6B0BFF5CB6F406B7ED" +
		"EE386BFB5A899FA5AE9F24117C4B1FE6" +
		"49286651ECE65381FFFFFFFFFFFFFFFF"
	Group2Generator = 2
)

func toString_DH_1024_BIT_MODP(attrType uint16, intValue uint16, bytesValue []byte) string {
	return DH_1024_BIT_MODP
}

var _ DHType = &Dh1024BitModp{}

type Dh1024BitModp struct {
	factor            *big.Int
	generator         *big.Int
	factorBytesLength int
}

func (t *Dh1024BitModp) TransformID() uint16 {
	return message.DH_1024_BIT_MODP
}

func (t *Dh1024BitModp) getAttribute() (bool, uint16, uint16, []byte) {
	return false, 0, 0, nil
}

func (t *Dh1024BitModp) GetSharedKey(secret, peerPublicValue *big.Int) []byte {
	sharedKey := new(big.Int).Exp(peerPublicValue, secret, t.factor).Bytes()
	prependZero := make([]byte, t.factorBytesLength-len(sharedKey))
	sharedKey = append(prependZero, sharedKey...)
	return sharedKey
}

func (t *Dh1024BitModp) GetPublicValue(secret *big.Int) []byte {
	localPublicValue := new(big.Int).Exp(t.generator, secret, t.factor).Bytes()
	prependZero := make([]byte, t.factorBytesLength-len(localPublicValue))
	localPublicValue = append(prependZero, localPublicValue...)
	return localPublicValue
}
