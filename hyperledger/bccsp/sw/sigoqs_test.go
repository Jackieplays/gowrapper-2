package sw

import (
	"testing"

	"github.com/hyperledger/fabric/sigoqs"
	"github.com/stretchr/testify/assert"
)

/*
import (
	"github.com/hyperledger/fabric/sigoqs"
	"testing"

	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/stretchr/testify/assert"
)
*/

/*
func TestSignECDSABadParameter(t *testing.T) {
	// Generate a key
	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	// Induce an error on the underlying ecdsa algorithm
	msg := []byte("hello world")
	oldN := lowLevelKey.Params().N
	defer func() { lowLevelKey.Params().N = oldN }()
	lowLevelKey.Params().N = big.NewInt(0)
	_, err = signECDSA(lowLevelKey, msg, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "zero parameter")
	lowLevelKey.Params().N = oldN
}

func TestVerifyECDSA(t *testing.T) {
	t.Parallel()

	// Generate a key
	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	msg := []byte("hello world")
	sigma, err := signECDSA(lowLevelKey, msg, nil)
	assert.NoError(t, err)

	valid, err := verifyECDSA(&lowLevelKey.PublicKey, sigma, msg, nil)
	assert.NoError(t, err)
	assert.True(t, valid)

	_, err = verifyECDSA(&lowLevelKey.PublicKey, nil, msg, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed unmashalling signature [")

	_, err = verifyECDSA(&lowLevelKey.PublicKey, nil, msg, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed unmashalling signature [")

	R, S, err := utils.UnmarshalECDSASignature(sigma)
	assert.NoError(t, err)
	S.Add(utils.GetCurveHalfOrdersAt(elliptic.P256()), big.NewInt(1))
	sigmaWrongS, err := utils.MarshalECDSASignature(R, S)
	assert.NoError(t, err)
	_, err = verifyECDSA(&lowLevelKey.PublicKey, sigmaWrongS, msg, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid S. Must be smaller than half the order [")
}
*/

func TestOQSSignerSign(t *testing.T) {
	t.Parallel()

	signer := &sigoqsSigner{}
	verifierPrivateKey := &sigoqsPrivateKeyVerifier{}
	verifierPublicKey := &sigoqsPublicKeyKeyVerifier{}

	lib, _ := sigoqs.LoadLib("/usr/local/lib/liboqs.so")
	defer lib.Close()

	sig, err := lib.GetSign(sigoqs.SigqTESLAI)

	sig.SetRandomAlg(AlgNistKat)

	// Get a particular KEM (don't forget to close)
	publicKey, secretKey, err := sig.KeyPair()
	defer sig.Close()

	assert.NoError(t, err)
	k := &sigoqsPrivateKey{&sigoqs.PrivateKey{sigoqs.PublicKey{publicKey}, secretKey}}
	pk, err := k.PublicKey()
	assert.NoError(t, err)

	// Sign
	msg := []byte("Hello World")
	sigma, err := signer.Sign(k, msg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, sigma)

	// Verify
	valid, err := verifyOQS(publicKey, msg, sigma)
	assert.NoError(t, err)
	assert.True(t, valid)

	valid, err = verifierPrivateKey.Verify(k, sigma, msg, nil)
	assert.NoError(t, err)
	assert.True(t, valid)

	valid, err = verifierPublicKey.Verify(pk, sigma, msg, nil)
	assert.NoError(t, err)
	assert.True(t, valid)
}

/*
func TestEcdsaPrivateKey(t *testing.T) {
	t.Parallel()

	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	k := &ecdsaPrivateKey{lowLevelKey}

	assert.False(t, k.Symmetric())
	assert.True(t, k.Private())

	_, err = k.Bytes()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Not supported.")

	k.privKey = nil
	ski := k.SKI()
	assert.Nil(t, ski)

	k.privKey = lowLevelKey
	ski = k.SKI()
	raw := elliptic.Marshal(k.privKey.Curve, k.privKey.PublicKey.X, k.privKey.PublicKey.Y)
	hash := sha256.New()
	hash.Write(raw)
	ski2 := hash.Sum(nil)
	assert.Equal(t, ski2, ski, "SKI is not computed in the right way.")

	pk, err := k.PublicKey()
	assert.NoError(t, err)
	assert.NotNil(t, pk)
	ecdsaPK, ok := pk.(*ecdsaPublicKey)
	assert.True(t, ok)
	assert.Equal(t, &lowLevelKey.PublicKey, ecdsaPK.pubKey)
}

func TestEcdsaPublicKey(t *testing.T) {
	t.Parallel()

	lowLevelKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	k := &ecdsaPublicKey{&lowLevelKey.PublicKey}

	assert.False(t, k.Symmetric())
	assert.False(t, k.Private())

	k.pubKey = nil
	ski := k.SKI()
	assert.Nil(t, ski)

	k.pubKey = &lowLevelKey.PublicKey
	ski = k.SKI()
	raw := elliptic.Marshal(k.pubKey.Curve, k.pubKey.X, k.pubKey.Y)
	hash := sha256.New()
	hash.Write(raw)
	ski2 := hash.Sum(nil)
	assert.Equal(t, ski, ski2, "SKI is not computed in the right way.")

	pk, err := k.PublicKey()
	assert.NoError(t, err)
	assert.Equal(t, k, pk)

	bytes, err := k.Bytes()
	assert.NoError(t, err)
	bytes2, err := x509.MarshalPKIXPublicKey(k.pubKey)
	assert.Equal(t, bytes2, bytes, "bytes are not computed in the right way.")

	invalidCurve := &elliptic.CurveParams{Name: "P-Invalid"}
	invalidCurve.BitSize = 1024
	k.pubKey = &ecdsa.PublicKey{Curve: invalidCurve, X: big.NewInt(1), Y: big.NewInt(1)}
	_, err = k.Bytes()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed marshalling key [")
}
*/
