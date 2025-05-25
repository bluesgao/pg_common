package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 加密（用公钥加密密码）
func RsaEncrypt(plainText []byte, publicKeyPEM []byte) (string, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return "", errors.New("无法解析公钥 PEM")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
	if err != nil {
		return "", err
	}

	// 返回 base64 编码后的密文，方便传输
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// 解密（用私钥解密密码）
func RsaDecrypt(base64CipherText string, privateKeyPEM []byte) (string, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return "", errors.New("无法解析私钥 PEM")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}
