package itp_request_client

import (
	"fmt"
	"github.com/franela/goreq"
	"github.com/jmcvetta/randutil"
	"github.com/pkg/errors"
	"github.com/yunmoon/itp-request-client/pkg/types"
	"github.com/yunmoon/itp-request-client/pkg/util"
	"github.com/yunmoon/itp-request-client/pkg/xrsa"
	"time"
)

const (
	Dev  = "dev"
	Test = "test"
	Prod = "prod"
)

var hosts = map[string]string{
	Dev:  "http://127.0.0.1:8092",
	Test: "",
	Prod: "",
}

type Options struct {
	AppId          string
	ThirdPublicKey string
	PrivateKey     string
	Version        string //版本号
	Env            string //环境
	Host           string //访问地址
}

func NewOptions(opt ...Option) Options {
	opts := Options{
		Version: "1.0",
		Env:     Test,
	}
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

type itpRequestClient struct {
	opts Options
}

func NewItpRequestClient(opt ...Option) (*itpRequestClient, error) {
	opts := NewOptions(opt...)
	if opts.AppId == "" {
		return nil, errors.New("appId未配置")
	}
	if opts.ThirdPublicKey == "" {
		return nil, errors.New("thirdPublicKey未配置")
	}
	if opts.PrivateKey == "" {
		return nil, errors.New("privateKey未配置")
	}
	return &itpRequestClient{opts: opts}, nil
}

func (client *itpRequestClient) Request(url string, body types.BodyMap, channelUserId string, encryptKeys []string, resp interface{}) error {
	var random string
	if len(encryptKeys) > 0 {
		randomStr, _ := randutil.AlphaString(16)
		for _, key := range encryptKeys {
			val := body.Get(key)
			if val != "" {
				encryptVal, err := util.AesCbcEncrypt(val, randomStr, randomStr)
				if err == nil {
					body.Set(key, encryptVal)
				}
			}
		}
		random, _ = xrsa.RsaEncrypt([]byte(randomStr), client.opts.ThirdPublicKey)
	}
	host := hosts[client.opts.Env]
	if client.opts.Host != "" {
		host = client.opts.Host
	}
	request := goreq.Request{
		Method:      "POST",
		Uri:         fmt.Sprintf("%s%s", host, url),
		Timeout:     5 * time.Second,
		ContentType: "application/json",
		Accept:      "application/json",
		Body:        body,
		Insecure:    true,
	}
	if channelUserId != "" {
		request.AddHeader("channelUserId", channelUserId)
	}
	if random != "" {
		request.AddHeader("random", random)
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	randomNum, _ := randutil.String(10, "0123456789")
	sequence := fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), randomNum)
	nonce, _ := randutil.AlphaString(32)
	signBody := make(types.BodyMap)
	signBody.Set("appid", client.opts.AppId).
		Set("version", client.opts.Version).
		Set("timestamp", timestamp).
		Set("sequence", sequence).
		Set("random", random).
		Set("noncestr", nonce).
		Set("channeluserid", channelUserId).
		Set("message", body.JsonBody())
	if client.opts.Env != Prod {
		fmt.Println("加签原始串：", signBody.EncodeURLParams())
	}
	sign, err := xrsa.GetRsaSign(signBody, xrsa.RSA2, client.opts.PrivateKey)
	if err != nil {
		return err
	}
	if client.opts.Env != Prod {
		fmt.Println("签名：", sign)
	}
	request.AddHeader("appid", client.opts.AppId)
	request.AddHeader("version", client.opts.Version)
	request.AddHeader("timestamp", timestamp)
	request.AddHeader("sequence", sequence)
	request.AddHeader("noncestr", nonce)
	request.AddHeader("sign", sign)
	res, err := request.Do()
	if err != nil {
		return err
	}
	return res.Body.FromJsonTo(resp)
}
