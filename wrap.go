	....................................................................
	func (s *sig) KeyPair() (publicKey, secretKey []byte, err error) {
		if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
		pubKeyLen := C.int(s.kem.length_public_key)
		pk := C.malloc(C.ulong(pubKeyLen))
		defer C.free(unsafe.Pointer(pk))
	
		secretKeyLen := C.int(k.kem.length_secret_key)
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
		//sig_len := len(signature) //Not sure...
 		defer C.free(unsafe.Pointer(sig1))

		messagelen := C.int(s.sig.length_message)
		mes := C.malloc(C.ulong(messagelen))
		//mes_len := len(message)
		defer C.free(unsafe.Pointer(mes))

        sk :=C.CBytes(secretKey) //Doubt...
		defer C.free(sk)
	

		res := C.sign(s.sig, (*C.uchar)(sig1), (*C.uint)(signaturelen), (*C.message)(mes),(*C.uint)(messagelen),(*C.uchar)(sk))
		if res != C.ERR_OK {
			return nil,libError(res, "signing failed")
		}
	
		return C.GoBytes(sig, signatureLen),  nil
	}

	-----------------------------------------------------Done------


func (s *sig) Verify(secretKey []byte,message []byte,signature []byte,publicKey []byte) (true or false []boolerr error) //Not sure
{
	if s.sig == nil {
			return nil, nil, errAlreadyClosed
		}
	
		signatureLen := C.int(s.sig.length_signature)
		sig1 := C.malloc(C.ulong(signatureLen))
		//sig_len := len(signature) //Not sure...
 		defer C.free(unsafe.Pointer(sig1))

		messagelen := C.int(s.sig.length_message)
		mes := C.malloc(C.ulong(messagelen))
		//mes_len := len(message)
		defer C.free(unsafe.Pointer(mes))

		pk :=C.CBytes(publicKey) //Doubt...
		defer C.free(pk)
	

        sk :=C.CBytes(secretKey) //Doubt...
		defer C.free(sk)
	

		res := C.sign(s.sig,(*C.message)(mes),(*C.uint)(messagelen),(*C.uchar)(sig1), (*C.uint)(signaturelen),(*C.uchar)(pk))
		if res != C.ERR_OK {
			return nil,libError(res, "verification failed")
		}
	
		return True,nil
	}
