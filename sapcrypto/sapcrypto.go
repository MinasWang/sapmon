package sapcrypto

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
)

const UKEY string = "9776Zhuyun1157Cloudcare6111Minas2869Wang5741400-882-33202178-jtcZY8V_XcCT20191020"

func AESEncrypt(src []byte ) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey([]byte(UKEY)))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AESDecrypt(encrypted []byte) ( []byte, error) {
	cipher, _ := aes.NewCipher(generateKey([]byte(UKEY)))
	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
		if trim > 0 {
			return decrypted[:trim], nil
		}
	}
	return decrypted, errors.New("faild to decrypted")
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func DecodeEncryption(str string) (string, error) {
	encrypted_v , _ := hex.DecodeString(str)
	decrypted, err := AESDecrypt(encrypted_v)
	if err != nil {
		// fmt.Println(err)
		return "", err
	} else {
		// fmt.Println(string(decrypted))
		return string(decrypted), nil
	}
}
