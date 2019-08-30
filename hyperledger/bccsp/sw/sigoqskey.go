/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sw

import (
	"crypto/sha256"

	"crypto/x509"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/sigoqs"
)

type sigoqsPrivateKey struct {
	privKey *sigoqs.PrivateKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *sigoqsPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k *sigoqsPrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}

	// Marshall the private key
	raw := k.privKey.Private_Key

	// Hash it
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *sigoqsPrivateKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *sigoqsPrivateKey) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *sigoqsPrivateKey) PublicKey() (bccsp.Key, error) {
	return &sigoqsPublicKey{&k.privKey.PublicKey}, nil
}

type sigoqsPublicKey struct {
	pubKey *sigoqs.PublicKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *sigoqsPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *sigoqsPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}

	// Marshall the public key
	raw := k.pubKey.Public_Key

	// Hash it
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *sigoqsPublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *sigoqsPublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *sigoqsPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}
