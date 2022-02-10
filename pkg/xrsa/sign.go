package xrsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"github.com/yunmoon/itp-request-client/v1/pkg/types"
	"github.com/yunmoon/itp-request-client/v1/pkg/xpem"
	"hash"
)

const (
	RSA  = "RSA"
	RSA2 = "RSA2"
)

func GetRsaSign(bm types.BodyMap, signType string, privateKeyStr string) (sign string, err error) {
	var (
		h              hash.Hash
		hashs          crypto.Hash
		encryptedBytes []byte
	)
	privateKeyStr = FormatPkcs1PrivateKey(privateKeyStr)
	privateKey, err := xpem.DecodePrivateKey([]byte(privateKeyStr))
	if err != nil {
		return
	}
	switch signType {
	case RSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	if _, err = h.Write([]byte(bm.EncodeURLParams())); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hashs, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

func VerifySign(signData, sign, signType, publicKeyStr string) (err error) {
	var (
		h     hash.Hash
		hashs crypto.Hash
	)
	publicKeyStr = FormatPublicKey(publicKeyStr)
	publicKey, err := xpem.DecodePublicKey([]byte(publicKeyStr))
	if err != nil {
		return err
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	switch signType {
	case RSA:
		hashs = crypto.SHA1
	case RSA2:
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}
