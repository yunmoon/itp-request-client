### Itp服务请求sdk

#### 如何使用

```golang
client, err := NewItpRequestClient(
		AppId(appId),
		Version(version),
		ThirdPublicKey(thirdPublicKey),
		PrivateKey(privateKey),
		Env(Dev),
		//Host("http://127.0.0.1"),
	)
if err == nil {
	respBody := make(types.BodyMap)
	reqBody := make(types.BodyMap)
	reqBody.Set("mobile", "18780101001")
	reqBody.Set("certType", 1)
	reqBody.Set("certNo", "18780101001")
	reqBody.Set("certName", "测试")
	err = client.Request("/v1/api/user/sync-user-info", reqBody, "test1", []string{"mobile", "certNo", "certName"}, &respBody)
	t.Logf("%v", respBody)
}
```