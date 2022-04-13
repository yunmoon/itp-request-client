package itp_request_client

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunmoon/itp-request-client/pkg/types"
	"testing"
)

var (
	appId          = "1234567812345678"
	version        = "1.0"
	thirdPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7sHkJQ1GrqJjsimN8BYifSM/aFDPcBPATnBnHaDbkK2weRtqi+cPPReWlRRFnsbbwxSPuLvNCQjRXtK4U2jk4UO0a/ZzehnzvJ/eT1CZP+g3o4RVJ3WO14ek4IpBku0pvdMfF1fN53dhwtSAqZ0q91GCRH/V6c8XVaWWE3fVOyXuezE6jm0DpfTY28VEaBiA+pA6INonKa29X4pgl7Ay6UpdTen6cg1kUmRn77uenTQC2QNcKWgUkV8I+sZ+eln7OYZ2S9pAwGuEiJKnzCqPsuZ0CJ4Tcz0qhFbW5eTxu8t1lppUdzQzZZFuAztEAeyPbIJ7SkMBTb5idup+O/zqFwIDAQAB"
	privateKey     = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDgz7ZO//H0xpbDKH+LDuBBm1qyc5BZYx46DSI8dJUM/v9iUjMdfE9/S3ip26qmc6XWWhRnRFMi/iTOHxvOLQ3OwUqXNQDsnVcFdLxxcDJaKFVPKu4rYQ5HtTpTD+Cl8wJcHb7LkDrU8t3nC/kQ7cH/lT/KDSjqaBBNeNceccMX0B9/wZWGu7M//CQqekcrJGhS6XDDtS7L97n/m63/aSFu/RCbM8q5RkwNPEjSDuf6l1q1yVa0fonewagAW2dSENAPC4UsBo6f84x97ViLbHXUaMujViuHXWyqOfGGvnZ+8IgcqJfcfH0lRxP7lp9xIi+u0eAXFwLnfT5VLNPAvIstAgMBAAECggEBAI9U3y9PD7y4QLb8wxStz10E57aO56GWCFeKuCFDUySOD9VoAx7xet32CGCDpTGq2jjoBcoxTbApyN2CCABTyVHg+uWc7ZHuXuUjoGHS3seMn5dyD5eosaoWabE7lkc3wHFqpZKzyk6q6b+9anbYn6+MQLdZ3JRW9M3wFXdboL80yGRMdbIBvEOOZFqS9xC9kWi6Ms0y8fxzcX3DbaqZKX2OT9Mc6aLLxup4ptCTPaXuHi/1HO1Be/cPHEBO0dHTJziznUvE3h46XHD//19jKx3nKXGp0enHHPvHjEmTAGCD4dalYrWfXBVor7NeZFgBDLMurp4XUn5ZBR1iRV01ngECgYEA9WIAtkymGiPYwaH9nSFdaw7T0YvPYfmrxViA9Yy+szgGUojeMsIEOt3WCIuxuHbeulJ0TKG/gTZ5TprLgmH/JZg2LedTUQhYtDvRuQCRU9cvoNjp2/qfZ5U4Mf8BJLe2d/oznqCpV3lIeGB/n8r5Whe342ihxUddkEGZKisyZH0CgYEA6onZOD7kJ1Do9fYWHe78vpZn5REkMyqWLOnUzjjirGvZ46gXj9TRLW2NQiE/cevaUqnkYBQHw1knvBC/xJWoIeM8Iq4dr/jL8L6gDoOJjuRC4DxDNxHbocMTo0qTDrSf2cM6Pxvz+d+HLbEPTiFY/Hs9qH7HrbwW1hTnaRW98HECgYEAxZSuAhvhuzaV+AQZlAYjlGqSAC5VRAynVPYYkJ9Nhj1cSeTPFYvHoCazipoA9gkw+lIeNv4el0pnjvVxXIDP01OmfHvBSIQx+J4aFp7wZdPlE9zVIT3CUMOERi2QnCIZGK4sFlRDRp3vzo3U9bOX6AUlGkVLzO/T1K4dSCkUIHkCgYBADUl/bN2ORzB4C67amev4eMcC7f1+48CDn5B4iVyOTh4BaGSW6T3/NA4B42aaTBkhvjgabR35oZ2SZNiabWyvZImFxxtgdYfsxYKctBubJIeHCa4pmfzrXoU8cR9cQsPtCr4bghzNPtiCB/rwEXdl7JpYK9eIgPeTm73fGwr2YQKBgD0hcjNRbhq0xIWqIwoC8yiEVaoZdvZoxXFFeyARKk2Lyjs30+k74LUmRhcBdiXAhXI8kIjgdsCquHRcETIE5b6diw5TXOSWoT/3Rl+eHmnnpFl1ES92LtZlv7SXEcvHrKRhem7YKvZQ0VA3dMqoiINdR32SLqgtK9c6KaiDkPw+"
)

func TestItpRequestClient_Request(t *testing.T) {
	client, err := NewItpRequestClient(
		AppId(appId),
		Version(version),
		ThirdPublicKey(thirdPublicKey),
		PrivateKey(privateKey),
		Env(Dev),
		Host("http://10.255.50.202:39092"),
	)
	if err == nil {
		respBody := make(types.BodyMap)
		reqBody := make(types.BodyMap)
		reqBody.Set("title", "测试推送").
			Set("content", "测试推送11111").
			Set("pushType", 1).
			Set("data", "{}")
		err = client.Request("/v1/api/user/sync-user-info", reqBody, "sdfghgfdsdfg", nil, &respBody)
		t.Logf("%v", respBody)
	}
	assert.NoError(t, err, "err:%v", err)
}
