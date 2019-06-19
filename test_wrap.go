  


package sigoliboqs

import (
	"testing"
  "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const libPath = "/usr/local/liboqs/lib/liboqs.so"

func TestRoundTrip(t *testing.T) {

	sigs := []SigTy
	SigPicnicL1FS   
     SigPicnicL1FS   
    SigPicnicL1FS    
   SigPicnicL1FS    
  SigPicnicL1FS    
   SigPicnicL1FS    
  SigPicnicL1FS    
  SigPicnicL1FS    
  SigPicnicL1FS    
  SigPicnicL1FS    
  SigPicnicL1FS    
  SigPicnicL1FS
		
	}

	k, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k.Close()) }()

	for _, sigAlg := range sigs {
		t.Run(string(sigAlg), func(t *testing.T) {
			//t.Parallel() <-- cannot use this because https://github.com/stretchr/testify/issues/187

			testSIG, err := k.GetSig(sigAlg)
			if err == errAlgDisabledOrUnknown {
				t.Skipf("Skipping disabled/unknown algorithm %q", sigAlg)
			}
			require.NoError(t, err)
			defer func() { require.NoError(t, testSIG.Close()) }()

			publicKey, secretKey, err := testSIG.KeyPair()
			require.NoError(t, err)

			sharedSecret, ciphertext, err := testSIG.Encaps(publicKey)
			require.NoError(t, err)

			recoveredSecret, err := testSIG.Decaps(ciphertext, secretKey)
			require.NoError(t, err)

			assert.Equal(t, sharedSecret, recoveredSecret)
		})
	}
}

func TestBadLibrary(t *testing.T) {
	_, err := LoadLib("bad")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to load module")
}

func TestReEntrantLibrary(t *testing.T) {
	k1, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k1.Close()) }()

	k2, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k2.Close()) }()
}

func TestLibraryClosed(t *testing.T) {
	k, err := LoadLib(libPath)
	require.NoError(t, err)
	require.NoError(t, k.Close())

	const expectedMsg = "library closed"

	t.Run("GetSIG", func(t *testing.T) {
		_, err := k.GetSig(KemBike1L1)
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})

	t.Run("Close", func(t *testing.T) {
		err := k.Close()
		require.Error(t, err)
		assert.Contains(t, err.Error(), expectedMsg)
	})
}

func TestKEMClosed(t *testing.T) {
	k, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k.Close()) }()

	testKEM, err := k.GetSig(KemKyber512)
	require.NoError(t, err)

	require.NoError(t, testSIG.Close())

	t.Run("KeyPair", func(t *testing.T) {
		_, _, err := testSIG.KeyPair()
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Encaps", func(t *testing.T) {
		_, _, err := testSIG.Encaps(nil)
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Decaps", func(t *testing.T) {
		_, err := testSIG.Decaps(nil, nil)
		assert.Equal(t, errAlreadyClosed, err)
	})

	t.Run("Decaps", func(t *testing.T) {
		err := testSIG.Close()
		assert.Equal(t, errAlreadyClosed, err)
	})
}

func TestInvalidSIGAlg(t *testing.T) {
	k, err := LoadLib(libPath)
	require.NoError(t, err)
	defer func() { require.NoError(t, k.Close()) }()

	_, err = k.GetSig(SigType("this will never be valid"))
	assert.Equal(t, errAlgDisabledOrUnknown, err)
}

func TestLibErr(t *testing.T) {
	// Difficult to test this without a deliberately failing KEM library (which could
	// be a future idea...)

	err := libError(operationFailed, "test%d", 123)
	assert.EqualError(t, err, "test123")
}



	

			

	

	

