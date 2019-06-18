package goliboqs

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -ldl
typedef enum {
	ERR_OK,
	ERR_CANNOT_LOAD_LIB,
	ERR_CONTEXT_CLOSED,
	ERR_MEM,
	ERR_NO_FUNCTION,
	ERR_OPERATION_FAILED,
} libResult;

#include <oqs/oqs.h>
#include <dlfcn.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  void *handle;
} ctx;

char *errorString(libResult r) {
	switch (r) {
	case ERR_CANNOT_LOAD_LIB:
		return "cannot load library";
	case ERR_CONTEXT_CLOSED:
		return "library closed";
	case ERR_MEM:
		return "out of memory";
	case ERR_NO_FUNCTION:
		return "library missing required function";
	case ERR_OPERATION_FAILED:
		// We have no further info to share
		return "operation failed";
	default:
		return "unknown error";
	}
}


libResult New(const char *path, ctx **c) {
	*c = malloc(sizeof(ctx));
	if (!(*c)) {
		return ERR_MEM;
	}
	(*c)->handle = dlopen(path, RTLD_NOW);
	if (NULL == (*c)->handle) {
		free(*c);
		return ERR_CANNOT_LOAD_LIB;
	}
	return ERR_OK;
}
libResult GetSign(const ctx *ctx, const char *name, OQS_SIG **sig) {
	if (!ctx->handle) {
		return ERR_CONTEXT_CLOSED;
	}
	// func matches signature of OQS_KEM_new
	OQS_SIG *(*func)(const char *);
	*(void **)(&func) = dlsym(ctx->handle, "OQS_SIG_new");
	if (NULL == func) {
		return ERR_NO_FUNCTION;
	}
	*sig = (*func)(name);
	return ERR_OK;
}




libResult FreeSig(ctx *ctx, OQS_SIG *sig) {
	if (!ctx->handle) {
		return ERR_CONTEXT_CLOSED;
	}
	void (*func)(OQS_SIG*);
	*(void **)(&func) = dlsym(ctx->handle, "OQS_SIG_free");
	if (NULL == func) {
		return ERR_NO_FUNCTION;
	}
	(*func)(sig);
	return ERR_OK;
}


libResult Close(ctx *ctx) {
	if (!ctx->handle) {
		return ERR_CONTEXT_CLOSED;
	}
	dlclose(ctx->handle);
	ctx->handle = NULL;
	return ERR_OK;
}


libResult KeyPair(const OQS_KEM *kem, uint8_t *public_key, uint8_t *secret_key) {
	OQS_STATUS status = sig->keypair(public_key, secret_key);
	if (status != OQS_SUCCESS) {
		return ERR_OPERATION_FAILED;
	}
	return ERR_OK;
}

libResult Sign(const OQS_SIG *sig, uint8_t *signature, size_t *signature_len, const uint8_t *message, size_t message_len, const uint8_t *secret_key) {
OQS_STATUS status =sig->sign(message, message_len, signature, signature_len, public_key) != OQS_SUCCESS) {
	if (status != OQS_SUCCESS) {
		return ERR_OPERATION_FAILED;
	}
	return ERR_OK;
}
libResult Verify(const OQS_SIG *sig, const uint8_t *message, size_t message_len, const uint8_t *signature, size_t signature_len, const uint8_t *public_key) {
	OQS_STATUS status =sig->verify(message, message_len, signature, signature_len, public_key) != OQS_SUCCESS) {
	if (status != OQS_SUCCESS) {
		return ERR_OPERATION_FAILED;
	}
	return ERR_OK;
}
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
---------------------------------------------------------Done-----------------------------------------------------------------
var errAlreadyClosed = errors.New("already closed")
var errAlgDisabledOrUnknown = errors.New("Signature algorithm is unknown or disabled")

var operationFailed C.libResult = C.ERR_OPERATION_FAILED


type sig struct {
	sig *C.OQS_SIG
	ctx *C.ctx
}


.......................................................Done...................................................................
	func (s *sig) KeyPair() (publicKey, secretKey []byte, err error) { //Not sure when to use []byte datatype ?
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
--------------------------------Done-------------------------------------------------------------------------------------
	func (s *sig) Sign(secretKey []byte,message []byte) (signature []byte, err error) {
		if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
		signatureLen := C.int(s.sig.length_signature)
		sig1 := C.malloc(C.ulong(signatureLen))
 		defer C.free(unsafe.Pointer(sig1))
		 
		mes_len := C.int(len(message))   //Should it be uint or int?, //Is len() fine?
		msg := C.CBytes(message) 
		defer C.free(msg)
	      
		

                sk :=C.CBytes(secretKey) 
		defer C.free(sk)
	

		res := C.Sign(s.sig, (*C.uchar)(sig1), (*C.int)(signatureLen), (*C.uchar)(msg),(*C.int)(mes_len),(*C.uchar)(sk))
		if res != C.ERR_OK {
			return nil,libError(res, "signing failed")
		}
	
		return C.GoBytes(sig1, signatureLen),  nil
	}

	-----------------------------------------------------Done-----------------------------------------------------




func (s *sig) Verify(message []byte,signature []byte,publicKey []byte) ([]bool ,err error) //Not sure
{
	if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
	mes_len := C.int(len(message))
	msg :=C.CBytes(message) 
		defer C.free(msg)
	
	
	sign_len := C.int(len(signature))
	sgn :=C.CBytes(signature) 
		defer C.free(sgn)
	

	pk :=C.CBytes(publicKey)
		defer C.free(pk)
	


		res := C.Verify(s.sig,(*C.uchar)(msg),(*C.int)(mes_len),(*C.uchar)(sgn), (*C.int)(sign_len),(*C.uchar)(pk))
		if res != C.ERR_OK {
			return nil,libError(res, "verification failed")
		}
	
		return true,nil
	}
-------------------------------------------------------------Done----------------------------------------------------------

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
-----------------------------------------------------Done--------------------------------------------------------------------
func libError(result C.libResult, msg string, a ...interface{}) error {
	
	if result == C.ERR_OPERATION_FAILED {
		return errors.Errorf(msg, a...)
	}

	str := C.GoString(C.errorString(result))
	return errors.Errorf("%s: %s", fmt.Sprintf(msg, a...), str)
}
-----------------------------------------------------------Done-------------------------------------------------------------
type Sig interface {
	
	KeyPair() (publicKey, secretKey []byte, err error)

	
	Sign(secretKey []byte,message []byte) (signature []byte, err error)

	
	Verify(message []byte,signature []byte,publicKey []byte) ([]bool ,err error)

	
	Close() error
}

---------------------------------------------------------------------Done---------------------------------------------------

type Lib struct {
	ctx *C.ctx
}
-----------------------------------------------------Done----------------------------------------------------------------
func (l *Lib) Close() error {
	res := C.Close(l.ctx)
	if res != C.ERR_OK {
		return libError(res, "failed to close library")
	}

	return nil
}
--------------------------------------------------------Done-----------------------------------------------------------------

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
-----------------------------------------------------------Done-----------------------------------------------------------------

func (l *Lib) GetSign(sigType sigType) (Sig, error) {
	cStr := C.CString(string(signType))
	defer C.free(unsafe.Pointer(cStr))

	var sigPtr *C.OQS_SIG

	res := C.GetSign(l.ctx, cStr, &sigPtr)
	if res != C.ERR_OK {
		return nil, libError(res, "failed to get Signature")
	}

	sig := &sig{
		sig: sigPtr,
		ctx: l.ctx,
	}
	if sig.sig == nil {
		return nil, errAlgDisabledOrUnknown
	}

	return sig, nil
}
----------------------------------------------------------------------Done------------------------------------------------------



