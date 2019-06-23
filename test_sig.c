#if defined(_WIN32)
#pragma warning(disable : 4244 4293)
#endif

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <oqs/oqs.h>

static OQS_STATUS sig_test_correctness(const char *method_name) {

	OQS_SIG *sig = NULL;
	uint8_t *public_key = NULL;
	uint8_t *secret_key = NULL;
	uint8_t *message = NULL;
	size_t message_len = 100;
	uint8_t *signature = NULL;
	size_t signature_len;
size_t secret_key_len;
size_t public_key_len;
	OQS_STATUS rc, ret = OQS_ERROR;

	sig = OQS_SIG_new(method_name);
	if (sig == NULL) {
		return OQS_SUCCESS;
	}

	printf("================================================================================\n");
	printf("Sample computation for signature %s\n", sig->method_name);
	printf("================================================================================\n");

	public_key = malloc(sig->length_public_key);
	secret_key = malloc(sig->length_secret_key);
	message = malloc(message_len);
	signature = malloc(sig->length_signature);
	signature_len = sig->length_signature;
        printf("See below \n");
         printf("Length of message: %zu \n",message_len);
          printf("Length of signature: %zu \n ",signature_len);
      
      secret_key_len = sig->length_secret_key;
      public_key_len =sig->length_public_key;

   




 printf("message:\n");
for (unsigned int i = 0; i < message_len; i++) {
  printf("%02X", message[i]);
}
printf("\n");
// printf("signature:\n");
/*
 printf("message:\n");
for (unsigned int i = 0; i < public_key; i++) {
  printf("%02X", public_key[i]);
}
printf("\n");
*/ 




printf("\n");


	if ((public_key == NULL) || (secret_key == NULL) || (message == NULL) || (signature == NULL)) {
		fprintf(stderr, "ERROR: malloc failed\n");
		goto err;
	}
    //printf("%u", (unsigned) public_key);
	OQS_randombytes(message, message_len);

	rc = OQS_SIG_keypair(sig, public_key, secret_key);
	if (rc != OQS_SUCCESS) {
		fprintf(stderr, "ERROR: OQS_SIG_keypair failed\n");
		goto err;
	}
 // printf("%p",(void*) sig->signature);


printf("\n");
printf("Length of secret-key: %zu \n ",secret_key_len);
printf("Length of public-key: %zu \n ",public_key_len);
 printf("Secret-key:\n");
for (unsigned int i = 0; i < secret_key_len; i++) {
  printf("%02X", secret_key[i]);
}
printf("\n");
printf("\n");

 printf("Public-key:\n");
for (unsigned int i = 0; i < public_key_len; i++) {
  printf("%02X", public_key[i]);
}
printf("\n");



// printf("signature:\n");

	rc = OQS_SIG_sign(sig, signature, &signature_len, message, message_len, secret_key);
	if (rc != OQS_SUCCESS) {
		fprintf(stderr, "ERROR: OQS_SIG_sign failed\n");
		goto err;
	}

printf("signature:\n");
for (unsigned int i = 0; i < signature_len; i++) {
  printf("%02X", signature[i]);
}
printf("\n");

	rc = OQS_SIG_verify(sig, message, message_len, signature, signature_len, public_key);
	if (rc != OQS_SUCCESS) {
		fprintf(stderr, "ERROR: OQS_SIG_verify failed\n");
		goto err;
	}

	/* modify the signature to invalidate it */
	signature[0]++;
	rc = OQS_SIG_verify(sig, message, message_len, signature, signature_len, public_key);
	if (rc != OQS_ERROR) {
		fprintf(stderr, "ERROR: OQS_SIG_verify should have failed!\n");
		goto err;
	}
	printf("verification passes as expected.Yes done!\n");
	ret = OQS_SUCCESS;
	goto cleanup;

err:
	ret = OQS_ERROR;

cleanup:
	if (sig != NULL) {
		OQS_MEM_secure_free(secret_key, sig->length_secret_key);
	}
	OQS_MEM_insecure_free(public_key);
	OQS_MEM_insecure_free(message);
	OQS_MEM_insecure_free(signature);
	OQS_SIG_free(sig);

	return ret;
}

int main() {
	int ret = EXIT_SUCCESS;
	OQS_STATUS rc;

	// Use system RNG in this program
	OQS_randombytes_switch_algorithm(OQS_RAND_alg_nist_kat);

	for (size_t i = 0; i < OQS_SIG_algs_length; i++) {
		rc = sig_test_correctness(OQS_SIG_alg_identifier(i));
		if (rc != OQS_SUCCESS) {
			ret = EXIT_FAILURE;
		}
	}

	return ret;
}
