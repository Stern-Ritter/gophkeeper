package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptData(t *testing.T) {
	key := []byte("HmQkWX1zTb6l3P8V8f3Eiw==")
	data := []byte("test data")

	_, err := EncryptData(data, key)
	assert.NoError(t, err, "expected no error encrypting data")
}

func TestDecryptData(t *testing.T) {
	key := []byte("HmQkWX1zTb6l3P8V8f3Eiw==")
	data := []byte("test data")

	encryptedData, err := EncryptData(data, key)
	assert.NoError(t, err, "expected no error encrypting data")

	decryptedData, err := DecryptData(encryptedData, key)
	assert.NoError(t, err, "expected no error decrypting data")

	assert.Equal(t, data, decryptedData, "decrypted data should match the original data")
}

func TestDecryptData_InvalidKey(t *testing.T) {
	key := []byte("HmQkWX1zTb6l3P8V8f3Eiw==")
	invalidKey := []byte("HmQkWX1z")

	data := []byte("test data")
	encryptedData, err := EncryptData(data, key)
	assert.NoError(t, err, "expected no error encrypting data")

	_, err = DecryptData(encryptedData, invalidKey)
	assert.Error(t, err, "expected error decrypting with an invalid key")
}

func TestDecryptData_InvalidData(t *testing.T) {
	key := []byte("HmQkWX1zTb6l3P8V8f3Eiw==")
	invalidData := []byte("invalid data")

	_, err := DecryptData(invalidData, key)
	assert.Error(t, err, "expected error while decrypting with invalid data")
}

func TestDecryptData_ShortData(t *testing.T) {
	key := []byte("HmQkWX1zTb6l3P8V8f3Eiw==")
	encryptedData := []byte("short")

	_, err := DecryptData(encryptedData, key)
	assert.Error(t, err, "expected error due to encrypted data being too short")
}

func TestEncryptData_InvalidKey(t *testing.T) {
	key := []byte("HmQkWX1z")
	data := []byte("test data")

	_, err := EncryptData(data, key)
	assert.Error(t, err, "expected error due to invalid key length")
}
