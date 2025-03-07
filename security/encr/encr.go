package encr

import (
	"github.com/pkg/errors"

	"github.com/nathaniel-bennett/ike/message"
	ikeCrypto "github.com/nathaniel-bennett/ike/security/IKECrypto"
)

var encrString map[uint16]func(uint16, uint16, []byte) string

var (
	encrTypes  map[string]ENCRType
	encrKTypes map[string]ENCRKType
)

func init() {
	// ENCR String
	encrString = make(map[uint16]func(uint16, uint16, []byte) string)
	encrString[message.ENCR_AES_CBC] = toString_ENCR_AES_CBC

	// ENCR Types
	encrTypes = make(map[string]ENCRType)

	encrTypes[ENCR_AES_CBC_128] = &EncrAesCbc{
		keyLength: 16,
	}
	encrTypes[ENCR_AES_CBC_192] = &EncrAesCbc{
		keyLength: 24,
	}
	encrTypes[ENCR_AES_CBC_256] = &EncrAesCbc{
		keyLength: 32,
	}

	// ENCR Kernel Types
	encrKTypes = make(map[string]ENCRKType)

	encrKTypes[ENCR_AES_CBC_128] = &EncrAesCbc{
		keyLength: 16,
	}
	encrKTypes[ENCR_AES_CBC_192] = &EncrAesCbc{
		keyLength: 24,
	}
	encrKTypes[ENCR_AES_CBC_256] = &EncrAesCbc{
		keyLength: 32,
	}
}

func StrToType(algo string) ENCRType {
	if t, ok := encrTypes[algo]; ok {
		return t
	} else {
		return nil
	}
}

func StrToKType(algo string) ENCRKType {
	if t, ok := encrKTypes[algo]; ok {
		return t
	} else {
		return nil
	}
}

func DecodeTransform(transform *message.Transform) ENCRType {
	if f, ok := encrString[transform.TransformID]; ok {
		s := f(transform.AttributeType, transform.AttributeValue, transform.VariableLengthAttributeValue)
		if s != "" {
			if encrType, ok2 := encrTypes[s]; ok2 {
				return encrType
			} else {
				return nil
			}
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func ToTransform(encrType ENCRType) (*message.Transform, error) {
	t := new(message.Transform)
	var err error
	t.TransformType = message.TypeEncryptionAlgorithm
	t.TransformID = encrType.TransformID()
	t.AttributePresent, t.AttributeType, t.AttributeValue, t.VariableLengthAttributeValue,
		err = encrType.getAttribute()
	if err != nil {
		return nil, errors.Wrapf(err, "ToTransform")
	}
	if t.AttributePresent && t.VariableLengthAttributeValue == nil {
		t.AttributeFormat = message.AttributeFormatUseTV
	}
	return t, nil
}

func DecodeTransformChildSA(transform *message.Transform) ENCRKType {
	if f, ok := encrString[transform.TransformID]; ok {
		s := f(transform.AttributeType, transform.AttributeValue, transform.VariableLengthAttributeValue)
		if s != "" {
			if encrKType, ok2 := encrKTypes[s]; ok2 {
				return encrKType
			} else {
				return nil
			}
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func ToTransformChildSA(encrKType ENCRKType) (*message.Transform, error) {
	t := new(message.Transform)
	var err error
	t.TransformType = message.TypeEncryptionAlgorithm
	t.TransformID = encrKType.TransformID()
	t.AttributePresent, t.AttributeType, t.AttributeValue, t.VariableLengthAttributeValue,
		err = encrKType.getAttribute()
	if err != nil {
		return nil, errors.Wrapf(err, "ToTransformChildSA")
	}
	if t.AttributePresent && t.VariableLengthAttributeValue == nil {
		t.AttributeFormat = 1 // TV
	}
	return t, nil
}

type ENCRType interface {
	TransformID() uint16
	getAttribute() (bool, uint16, uint16, []byte, error)
	GetKeyLength() int
	NewCrypto(key []byte) (ikeCrypto.IKECrypto, error)
}

type ENCRKType interface {
	TransformID() uint16
	getAttribute() (bool, uint16, uint16, []byte, error)
	GetKeyLength() int
}
