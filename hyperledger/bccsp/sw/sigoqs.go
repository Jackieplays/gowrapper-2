package sw

import (
	"errors"

	"github.com/hyperledger/fabric/sigoqs"

	"github.com/hyperledger/fabric/bccsp"
)

//func signOQS(secretKey []byte, digest []byte, params *sig) ([]byte, error) {
func signOQS(secretKey []byte, digest []byte) ([]byte, error) {
	lib, _ := sigoqs.LoadLib("/usr/local/lib/liboqs.so")
	defer lib.Close()

	sig, err := lib.GetSign(sigoqs.SigqTESLAI)
	if err != nil {
		return nil, err
	}
	s, err2 := sig.Sign(secretKey, digest)
	defer sig.Close()

	if err2 != nil {
		return nil, err2
	}

	return s, nil
}

func verifyOQS(publicKey, digest, signature []byte) (bool, error) {

	lib, _ := sigoqs.LoadLib("/usr/local/lib/liboqs.so")
	defer lib.Close()

	sig, err := lib.GetSign(sigoqs.SigqTESLAI)
	if err != nil {
		return false, err
	}
	defer sig.Close()
	return sig.Verify(digest, signature, publicKey)
}

type sigoqsSigner struct{}

func (s *sigoqsSigner) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	if opts == nil {
		return nil, errors.New("Invalid options. Must be different from nil.")
	}

	ss, err := signOQS(k.(*sigoqsPrivateKey).privKey.Private_Key, digest)
	if err != nil {
		return nil, err
	}

	return ss, nil
}

type sigoqsPrivateKeyVerifier struct{}

func (v *sigoqsPrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	if opts == nil {
		return false, errors.New("Invalid options. It must not be nil.")
	}

	return verifyOQS(k.(*sigoqsPrivateKey).privKey.Public_Key, digest, signature)
}

type sigoqsPublicKeyKeyVerifier struct{}

func (v *sigoqsPublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	if opts == nil {
		return false, errors.New("Invalid options. It must not be nil.")
	}

	return verifyOQS(k.(*sigoqsPublicKey).pubKey.Public_Key, digest, signature)
}
