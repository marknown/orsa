package orsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// EncryptString use a public key encrypt string
func EncryptString(text string, publicKey string, isBase64 bool) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	cipher, err2 := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(text))
	if err2 != nil {
		return "", err2
	}

	if isBase64 {
		return base64.StdEncoding.EncodeToString(cipher), nil
	} else {
		return string(cipher), nil
	}
}

// DecryptString use a private key decrypt string
func DecryptString(ciphertext string, privateKey string, isBase64 bool) (string, error) {
	var cipher []byte
	var err error
	if isBase64 {
		cipher, err = base64.StdEncoding.DecodeString(ciphertext)
		if nil != err {
			return "", err
		}
	} else {
		cipher = []byte(ciphertext)
	}

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err2 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if nil != err2 {
		return "", err2
	}

	textBtyes, err3 := rsa.DecryptPKCS1v15(rand.Reader, priv, cipher)

	return string(textBtyes), err3
}

// StringToPubKey string to public key
func StringToPubKey(s string) string {
	return `
-----BEGIN PUBLIC KEY-----
` + s + `
-----END PUBLIC KEY-----
`
}

// StringToPrivKey string to private key
func StringToPrivKey(s string) string {
	return `
-----BEGIN RSA PRIVATE KEY-----
` + s + `
-----END RSA PRIVATE KEY-----
`
}
