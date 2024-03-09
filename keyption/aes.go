package keyption

//加解密相关
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func AESencrypt(plainText string, keyString string) (string, error) {
	// Convert the plaintext and key to bytes
	plaintext := []byte(plainText)
	key := []byte(keyString)
	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plaintext = ZeroPadding([]byte(plaintext), block.BlockSize())
	iv := FillIV(keyString) // 初始化向量（IV）需要与解密时的IV相同
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	// Return the encrypted ciphertext as a hex-encoded string
	return hex.EncodeToString(ciphertext), nil
}

func AESdecrypt(ciphertextString string, keyString string) (string, error) {
	// Convert the ciphertext and key from hex-encoded strings to bytes
	ciphertext, err := hex.DecodeString(ciphertextString)
	if err != nil {
		return "", err
	}
	key := []byte(keyString)
	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := FillIV(keyString) // 初始化向量（IV）需要与解密时的IV相同
	plaintext := make([]byte, len(ciphertext))
	mode2 := cipher.NewCBCDecrypter(block, iv)
	mode2.CryptBlocks(plaintext, ciphertext)
	// 去除ZeroPadding填充数据
	plaintext = ZeroUnPadding(plaintext)
	// Convert the plaintext to a string and return it
	return string(plaintext), nil

}

// ZeroPadding使用0填充数据
func ZeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padData := bytes.Repeat([]byte{0}, padding)
	return append(data, padData...)
}

// ZeroUnPadding去除ZeroPadding填充的数据
func ZeroUnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// ZeroUnPadding去除ZeroPadding填充的数据
func FillIV(keyString string) []byte {
	key := []byte(keyString)[:16]
	if len(key) < 16 {
		for i := len(key); i < 16; i++ {
			key = append(key, []byte("=")[0])
		}
	}
	return key
}
