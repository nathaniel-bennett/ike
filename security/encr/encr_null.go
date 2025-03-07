package encr

import (
	"github.com/nathaniel-bennett/ike/message"
	ikeCrypto "github.com/nathaniel-bennett/ike/security/IKECrypto"
)

const (
	ENCR_NULL string = "ENCR_NULL"
)

func toString_ENCR_NULL(attrType uint16, intValue uint16, bytesValue []byte) string {
	return ENCR_NULL
}

var (
	_ ENCRType  = &EncrNull{}
	_ ENCRKType = &EncrNull{}
)

type EncrNull struct {
}

func (t *EncrNull) TransformID() uint16 {
	return message.ENCR_NULL
}

func (t *EncrNull) getAttribute() (bool, uint16, uint16, []byte, error) {
	return true, message.AttributeTypeKeyLength, uint16(0), nil, nil
}

func (t *EncrNull) GetKeyLength() int {
	return 0
}

func (t *EncrNull) NewCrypto(key []byte) (ikeCrypto.IKECrypto, error) {
	encr := new(EncrNullCrypto)
	return encr, nil
}

var _ ikeCrypto.IKECrypto = &EncrNullCrypto{}

type EncrNullCrypto struct {
}

func (encr *EncrNullCrypto) Encrypt(plainText []byte) ([]byte, error) {
	return plainText, nil
}

func (encr *EncrNullCrypto) Decrypt(cipherText []byte) ([]byte, error) {
	return cipherText, nil
}
