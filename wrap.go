	





*/
/////////////////////////////////////////////////////////////////////////
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
)

type SigType string

const (

     SigPicnicL1FS   SigType="picnic_L1_FS"
     SigPicnicL1FS   SigType="picnic_L1_UR"

    SigPicnicL1FS    SigType="picnic_L3_FS"

  SigPicnicL1FS    SigType="picnic_L3_UR"

  SigPicnicL1FS    SigType="picnic_L5_FS"

   SigPicnicL1FS    SigType="picnic_L5_UR"

  SigPicnicL1FS    SigType="picnic2_L1_FS"

  SigPicnicL1FS    SigType="picnic2_L3_FS"

  SigPicnicL1FS    SigType="picnic2_L5_FS"

  SigPicnicL1FS    SigType="qTESLA_I"

  SigPicnicL1FS    SigType="qTESLA_III_size"

  SigPicnicL1FS    SigType="qTESLA_III_speed"

)
---------------------------------------------------------------------------------
var errAlreadyClosed = errors.New("already closed")
var errAlgDisabledOrUnknown = errors.New("Signature algorithm is unknown or disabled")

// operationFailed exposed to help test code (which cannot use cgo "C.<foo>" variables)
var operationFailed C.libResult = C.ERR_OPERATION_FAILED


type sign struct {
	sign *C.OQS_SIG
	ctx *C.ctx
}


....................................................................
	func (s *sig) KeyPair() (publicKey, secretKey []byte, err error) {
		if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
		pubKeyLen := C.int(s.sig.length_public_key)
		pk := C.malloc(C.ulong(pubKeyLen))
		defer C.free(unsafe.Pointer(pk))
	
		secretKeyLen := C.int(k.sig.length_secret_key)
		sk := C.malloc(C.ulong(secretKeyLen))
		defer C.free(unsafe.Pointer(sk))
	
		res := C.KeyPair(s.sig, (*C.uchar)(pk), (*C.uchar)(sk))
		if res != C.ERR_OK {
			return nil, nil, libError(res, "key pair generation failed")
		}
	
		return C.GoBytes(pk, pubKeyLen), C.GoBytes(sk, secretKeyLen), nil
	}
--------------------------------Done----------------------------------------------------
	func (s *sig) Sign(secretKey []byte,message []byte) (signature []byte, err error) {
		if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
		signatureLen := C.int(s.sig.length_signature)
		sig1 := C.malloc(C.ulong(signatureLen))
 		defer C.free(unsafe.Pointer(sig1))
		 
		msg := C.CBytes(message) 
		defer C.free(msg)
	        mes_len := len(message)
		

                sk :=C.CBytes(secretKey) 
		defer C.free(sk)
	

		res := C.sign(s.sig, (*C.uchar)(sig1), (*C.uint)(signaturelen), (*C.message)(mes),(*C.uint)(mes_len),(*C.uchar)(sk))
		if res != C.ERR_OK {
			return nil,libError(res, "signing failed")
		}
	
		return C.GoBytes(sig, signatureLen),  nil
	}

	-----------------------------------------------------Done------




func (s *sig) Verify(secretKey []byte,message []byte,signature []byte,publicKey []byte) ([]bool ,error) //Not sure
{
	if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	sk :=C.CBytes(secretKey) 
		defer C.free(sk)
	
	mes_len := C.uint(len(message))
	msg :=C.CBytes(message) 
		defer C.free(msg)
	
	
	sign_len := C.uint(len(signature))
	sgn :=C.CBytes(signature) 
		defer C.free(sgn)
	

	pk :=C.CBytes(publicKey)
		defer C.free(pk)
	


		res := C.Verify(s.sig,(*C.message)(msg),(*C.uint)(mes_len),(*C.uchar)(sgn), (*C.uint)(sign_len),(*C.uchar)(pk))
		if res != C.ERR_OK {
			return nil,libError(res, "verification failed")
		}
	
		return true,nil
	}
-------------------------------------------------------------Done------------------------------------------------------

func (s *sig) Close() error {
	if s.sig == nil {
		return errAlreadyClosed
	}

	res := C.FreeSig(k.ctx, s.sig)
	if res != C.ERR_OK {
		return libError(res, "failed to free signature")
	}

	s.sig = nil
	return nil
}
-----------------------------------------------------Done------------------------
func libError(result C.libResult, msg string, a ...interface{}) error {
	
	if result == C.ERR_OPERATION_FAILED {
		return errors.Errorf(msg, a...)
	}

	str := C.GoString(C.errorString(result))
	return errors.Errorf("%s: %s", fmt.Sprintf(msg, a...), str)
}
-----------------------------------------------------------Done----------------------------------------
type Kem interface {
	// KeyPair generates a new key pair.
	KeyPair() (publicKey, secretKey []byte, err error)

	// Encaps generates a new shared secret and encrypts it under the public key.
	
	Sign(secretKey []byte,message []byte) (signature []byte, err error)

	
	Verify(secretKey []byte,message []byte,signature []byte,publicKey []byte) ([]bool ,error)

	
	Close() error
}

---------------------------------------------------------------------Done---------------------------------------------

type Lib struct {
	ctx *C.ctx
}
-----------------------------------------------------Done--------------------------------
func (l *Lib) Close() error {
	res := C.Close(l.ctx)
	if res != C.ERR_OK {
		return libError(res, "failed to close library")
	}

	return nil
}
--------------------------------------------------------Done----------------------------

// LoadLib loads the liboqs library. The path parameter is given directly to dlopen, see the dlopen man page
// for details of how path is interpreted. (Paths with a slash are treated as absolute or relative paths). Be
// sure to Close after use to free resources.
func LoadLib(path string) (*Lib, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	var ctx *C.ctx
	res := C.New(p, &ctx)
	if res != C.ERR_OK {
		return nil, libError(res, "failed to load module at %q", path)
	}

	return &Lib{ctx: ctx}, nil
}
-----------------------------------------------------------Done------------------------------
/ GetKem returns a Kem for the specified algorithm. Constants are provided for known algorithms,
// but any string can be provided and will be passed through to liboqs. As a reminder, some algorithms
// need to be explicitly enabled when building liboqs.
func (l *Lib) GetSign(signType signType) (Sign, error) {
	cStr := C.CString(string(signType))
	defer C.free(unsafe.Pointer(cStr))

	var kemPtr *C.OQS_SIG

	res := C.GetSign(l.ctx, cStr, &signPtr)
	if res != C.ERR_OK {
		return nil, libError(res, "failed to get Signature")
	}

	sign := &sign{
		sign: signPtr,
		ctx: l.ctx,
	}
	if sign.sign == nil {
		return nil, errAlgDisabledOrUnknown
	}

	return sign, nil
}
----------------------------------------------------------------------Done--------------------------------------------



