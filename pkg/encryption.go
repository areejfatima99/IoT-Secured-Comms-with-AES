package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type Encryption struct {
	Key []byte
}

func GenerateRandomKey() []byte {
	key := make([]byte, 32) // AES-256 key size
	_, err := rand.Read(key)
	if err != nil {
		panic("Failed to generate random key")
	}
	return key
}

func (e *Encryption) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(generateChebyshevKey(aes.BlockSize))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	fmt.Printf("plaintext: %s\n", plaintext)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func (e *Encryption) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(generateChebyshevKey(aes.BlockSize))
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	shiftParam := len(plaintext) % aes.BlockSize
	plaintext = dynamicShiftRows(plaintext, aes.BlockSize-shiftParam)

	return plaintext, nil
}

// Dynamic shift rows based on a dynamic parameter
func dynamicShiftRows(data []byte, shiftParam int) []byte {
	blockSize := aes.BlockSize
	numRows := len(data) / blockSize
	shiftedData := make([]byte, len(data))

	for row := 0; row < numRows; row++ {
		shiftAmount := shiftParam % blockSize
		startIdx := row * blockSize
		endIdx := (row + 1) * blockSize

		copy(shiftedData[startIdx:endIdx], data[startIdx+shiftAmount:endIdx])
		copy(shiftedData[startIdx+blockSize-shiftAmount:endIdx], data[startIdx:endIdx-shiftAmount])
	}

	return shiftedData
}

// Implement Chebyshev map iteration to generate a key
func generateChebyshevKey(keySize int) []byte {
	// Initialize seed and other parameters
	seed := uint32(123456789) // Initial seed
	numIterations := 100      // Number of iterations
	key := make([]byte, keySize)

	// Iterate the Chebyshev map
	for i := 0; i < numIterations; i++ {
		// Update the seed using Chebyshev map iteration
		// This involves a mathematical function involving the previous seed
		seed = (seed * seed) % 255

		// Use the updated seed to generate key bytes
		key[i%keySize] = byte(seed)
	}

	return key
}
