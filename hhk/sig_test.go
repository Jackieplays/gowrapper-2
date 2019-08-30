package sigoliboqs

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const libPath = "/usr/local/lib/liboqs.so"

func TestRoundTrip(t *testing.T) {

	sigs := []SigType{
		SigqTESLAIIIspeed,
		SigqTESLAI,
		SigqTESLAIIIsize,

		//	SigPicnicL1FS,
		//SigPicnicL1UR,
		//SigPicnicL3FS,
		//SigPicnicL3UR,
		//SigPicnicL5FS,
		//SigPicnicL5UR,
		//SigPicnic2L1FS,
		//SigPicnic2L3FS,
		//SigPicnic2L5FS,
	}

	//var b string
	s, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, s.Close()) }()

	//message := make([]byte, 100)
	s.SetRandomAlg(AlgNistKat)
	//s.SetRandomBytes(message, 100)
	message, err := s.GetRandomBytes(100)
	fmt.Println("message")
	h12 := strings.ToUpper(hex.EncodeToString(message))
	fmt.Printf("%s\n", h12)

	for _, sigAlg := range sigs {
		t.Run(string(sigAlg), func(t *testing.T) {
			//s.SetRandomAlg(AlgNistKat)

			testSIG, err := s.GetSign(sigAlg)
			if err == errAlgDisabledOrUnknown {
				t.Skipf("Skipping disabled/unknown algorithm %q", sigAlg)
			}
			require.NoError(t, err)
			defer func() { require.NoError(t, testSIG.Close()) }()
			//fmt.Println("message")
			start := time.Now()
			publicKey, secretKey, err := testSIG.KeyPair()
			elapsed := time.Since(start)
			fmt.Printf("1. b) key-pair gen1  \n%s", elapsed)
			require.NoError(t, err)

			fmt.Println("GeneratedSecretkey")
			h := strings.ToUpper(hex.EncodeToString(secretKey))
			fmt.Printf("%s\n", h)

			fmt.Println("GeneratedPublickey")
			h1 := strings.ToUpper(hex.EncodeToString(publicKey))
			fmt.Printf("%s\n", h1)

			//I think the fuction below takes the generatedsecretKEy and not the pre-defined one as the signature changes everytime I run.But the o/p in (***)is the pre-defined one. It's contradicting.
			signature, err := testSIG.Sign(secretKey, message)
			require.NoError(t, err)

			fmt.Println("GeneratedSignature")
			h2 := strings.ToUpper(hex.EncodeToString(signature))
			fmt.Printf("%s\n", h2)
			//fmt.Println("publickey")
			//fmt.Println("signature")
			//h3 := hex.EncodeToString(signature)
			//	fmt.Println(h3)

			result, err := testSIG.Verify(message, signature, publicKey) //assert is of type bool
			require.NoError(t, err)
			///fmt.Println("message")
			//b := string(secretKey[:])
			//fmt.Println(b)
			//b := string(secretKey[:])
			//fmt.Println(secretKey)
			//hex.EncodeToString()

			assert.Equal(t, result, true)
		})
	}

}

func TestBadLibrary(t *testing.T) {
	//	fmt.Println("message-6")
	_, err := LoadLib("bad")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to load module")
}

func TestReEntrantLibrary(t *testing.T) {
	//fmt.Println("message-5")
	s1, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, s1.Close()) }()

	s2, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, s2.Close()) }()
}

func TestLibraryClosed(t *testing.T) {
	//fmt.Println("message-4")
	s, err := LoadLib(libPath)
	require.NoError(t, err)
	require.NoError(t, s.Close())

	const expectedMsg = "library closed"

	t.Run("GetSIG", func(t *testing.T) {
		//fmt.Println("message-3")
		_, err := s.GetSign(SigPicnicL1FS)
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("Close", func(t *testing.T) {
		//fmt.Println("message-2")
		err := s.Close()
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})
}

func TestSIGClosed(t *testing.T) {
	//fmt.Println("message-1")
	s, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, s.Close()) }()

	testSIG, err := s.GetSign(SigqTESLAI)
	require.NoError(t, err)

	require.NoError(t, testSIG.Close())

	t.Run("KeyPair", func(t *testing.T) {
		fmt.Println("message0")
		_, _, err := testSIG.KeyPair()
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Sign", func(t *testing.T) {
		fmt.Println("message1")
		_, err := testSIG.Sign(nil, nil)
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Verify", func(t *testing.T) {
		fmt.Println("message2")
		_, err := testSIG.Verify(nil, nil, nil)
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Verify", func(t *testing.T) {
		fmt.Println("message3")
		err := testSIG.Close()
		assert.Equal(t, errAlreadyClosed, err)
	})
}

func TestInvalidSIGAlg(t *testing.T) {
	fmt.Println("message4")
	s, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, s.Close()) }()

	_, err = s.GetSign(SigType("this will never be valid"))
	assert.Equal(t, errAlgDisabledOrUnknown, err)
}

func TestLibErr(t *testing.T) {
	fmt.Println("message5")
	err := libError(operationFailed, "test%d", 123)
	assert.EqualError(t, err, "test123")
}
