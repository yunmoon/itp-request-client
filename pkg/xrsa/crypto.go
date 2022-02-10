package xrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"github.com/yunmoon/itp-request-client/pkg/xpem"
)

func RsaEncrypt(str []byte, publicKeyStr string) (string, error) {
	publicKeyStr = FormatPublicKey(publicKeyStr)
	publicKey, err := xpem.DecodePublicKey([]byte(publicKeyStr))
	if err != nil {
		return "", err
	}
	//rsa公钥加密
	result, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, str)
	if err != nil {
		return "", err
	}
	encodeStr := base64.StdEncoding.EncodeToString(result)
	return encodeStr, nil
}

func RsaDecrypt(str string, privateKeyStr string) (string, error) {
	privateKeyStr = FormatPkcs1PrivateKey(privateKeyStr)
	privateKey, err := xpem.DecodePrivateKey([]byte(privateKeyStr))
	if err != nil {
		return "", err
	}
	//base64解密
	decodeData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	//rsa私钥解密
	decodeData, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeData)
	if err != nil {
		return "", err
	}
	return string(decodeData), nil
}
