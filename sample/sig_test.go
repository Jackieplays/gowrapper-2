package sigoliboqs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const libPath = "/usr/local/liboqs/.libs/liboqs.so"

func TestRoundTrip(t *testing.T) {

	sigs := []SigType{
		//SigPicnicL1FS,
		SigPicnicL1UR,
		//SigPicnicL3FS,
		//SigPicnicL3UR,
		//SigPicnicL5FS,
		//SigPicnicL5UR,
		SigPicnic2L1FS,
		//SigPicnic2L3FS,
		//SigPicnic2L5FS,
		//SigqTESLAI,
		//SigqTESLAIIIsize,
		//SigqTESLAIIIspeed,
	}

	var message = []byte("Hello")
	//var b string
	k, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k.Close()) }()

	for _, sigAlg := range sigs {
		t.Run(string(sigAlg), func(t *testing.T) {

			testSIG, err := k.GetSign(sigAlg)
			if err == errAlgDisabledOrUnknown {
				t.Skipf("Skipping disabled/unknown algorithm %q", sigAlg)
			}
			require.NoError(t, err)
			defer func() { require.NoError(t, testSIG.Close()) }()
			fmt.Println("message")
			publicKey, secretKey, err := testSIG.KeyPair()
			//b := []byte{publicKey}
			//s := string(b)
			/*
				b := string(secretKey[:])
				fmt.Println(b)
			*/
			//fmt.Println(secretKey)

			require.NoError(t, err)
			fmt.Println("message")
			signature, err := testSIG.Sign(secretKey, message)
			require.NoError(t, err)
			fmt.Println("message")
			result, err := testSIG.Verify(message, signature, publicKey) //assert is of type bool
			require.NoError(t, err)
			fmt.Println("message")

			assert.Equal(t, result, true)
		})
	}
}

func TestBadLibrary(t *testing.T) {
	fmt.Println("message-6")
	_, err := LoadLib("bad")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to load module")
}

func TestReEntrantLibrary(t *testing.T) {
	fmt.Println("message-5")
	k1, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k1.Close()) }()

	k2, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k2.Close()) }()
}

func TestLibraryClosed(t *testing.T) {
	fmt.Println("message-4")
	s, err := LoadLib(libPath)
	require.NoError(t, err)
	require.NoError(t, s.Close())

	const expectedMsg = "library closed"

	t.Run("GetSIG", func(t *testing.T) {
		fmt.Println("message-3")
		_, err := s.GetSign(SigPicnicL1FS)
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("Close", func(t *testing.T) {
		fmt.Println("message-2")
		err := s.Close()
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})
}

func TestSIGClosed(t *testing.T) {
	fmt.Println("message-1")
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
