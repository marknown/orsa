# orsa
orsa is a go rsa encrypt/decrypt package

## EncryptString use a public key encrypt string
```
func EncryptString(text string, publicKey string, isBase64 bool) (string, error) {
}

Usage:

orsa.EncryptString("aaa", "public key string", true)
```

### DecryptString use a private key decrypt string
```
func DecryptString(ciphertext string, privateKey string, isBase64 bool) (string, error) {
}

Usage:

orsa.DecryptString("aaa", "privateKey string", true)
```

### StringToPubKey && StringToPrivKey
```
If a public key is not contains "-----BEGIN PUBLIC KEY-----"
you must use StringToPubKey function to combine a public key


func StringToPubKey(s string) string {
    return `
-----BEGIN PUBLIC KEY-----
` + s + `
-----END PUBLIC KEY-----
`
}


If a private key is not contains "-----BEGIN RSA PRIVATE KEY-----"
you must use StringToPubKey function to combine a private key


func StringToPrivKey(s string) string {
    return `
-----BEGIN RSA PRIVATE KEY-----
` + s + `
-----END RSA PRIVATE KEY-----
`
}
```