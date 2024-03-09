package keyption

//加解密相关
import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
)

var RSA_PUBLIC_KEY *rsa.PublicKey
var RSA_PRIVATE_KEY *rsa.PrivateKey

func ReadPrivateKeyFromText(keyString string) (*rsa.PrivateKey, error) {
	privateKeyFile := []byte(keyString)
	block, _ := pem.Decode(privateKeyFile)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func ReadPublicKeyFromText(keyString string) (*rsa.PublicKey, error) {
	publicKeyFile := []byte(keyString)
	block, _ := pem.Decode(publicKeyFile)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse RSA public key")
	}
	return publicKey, nil
}

func ReplacePublicKeyFromText(keyString string) string {
	replacedStr := keyString
	replacedStr = strings.Replace(replacedStr, "-----BEGIN RSA PUBLIC KEY-----", "", 1)
	replacedStr = strings.Replace(replacedStr, "-----END RSA PUBLIC KEY-----", "", 1)
	replacedStr = strings.Replace(replacedStr, " ", "", -1)
	replacedStr = strings.TrimSpace(replacedStr)
	return "-----BEGIN RSA PUBLIC KEY-----\n" + replacedStr + "\n-----END RSA PUBLIC KEY-----"
}

/*
RSA加密时，对于原文数据的要求：
OAEP填充模式： 原文长度 <= 密钥模长 - (2 * 原文的摘要值长度) - 2字节
        各摘要值长度：
                SHA-1:    20字节
                SHA-256:  32字节
                SHA-384:  48字节
                SHA-512:  64字节
PKCA1-V1_5填充模式：原文长度 <= 密钥模长 - 11字节
*/

func RSApublicEncrypt(rsaPub *rsa.PublicKey, ciphertextString string) (string, error) {
	plaintext := []byte(ciphertextString)
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPub, plaintext, nil)
	if err != nil {
		log.Println("RSA encryption failed:", err)
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func RSAprivateDecrypt(rsaPri *rsa.PrivateKey, ciphertextString string) (string, error) {
	plaintext, err := hex.DecodeString(ciphertextString)
	if err != nil {
		return "", err
	}
	decryptedText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPri, plaintext, nil)
	if err != nil {
		log.Println("RSA decryption failed:", err)
		return "", err
	}
	return string(decryptedText), nil
}

func RSApublicEncryptBlock(rsaPub *rsa.PublicKey, ciphertextString string) (string, error) {
	plaintext := []byte(ciphertextString)

	keySize, srcSize := rsaPub.Size(), len(plaintext)
	//单次加密的长度需要减掉padding的长度
	padding := 2*32 + 2
	offSet, once := 0, keySize-padding
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPub, plaintext[offSet:endIndex], nil)
		if err != nil {
			log.Println("RSA encryption block failed:", err)
			return "", err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt := buffer.Bytes()
	return hex.EncodeToString(bytesEncrypt), nil
}

func RSAprivateDecryptBlock(rsaPri *rsa.PrivateKey, ciphertextString string) (string, error) {
	encryptedData, err := hex.DecodeString(ciphertextString)
	if err != nil {
		return "", err
	}
	keySize, srcSize := rsaPri.Size(), len(encryptedData)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPri, encryptedData[offSet:endIndex], nil)
		if err != nil {
			log.Println("RSA decrypt block failed:", err)
			return "", err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt := buffer.Bytes()
	plaintext := string(bytesDecrypt)
	return plaintext, nil
}
