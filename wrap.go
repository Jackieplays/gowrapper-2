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
