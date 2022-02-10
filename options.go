package itp_request_client

type Option func(*Options)

func AppId(appId string) Option {
	return func(o *Options) {
		o.AppId = appId
	}
}

func Env(env string) Option {
	return func(options *Options) {
		options.Env = env
	}
}

func Host(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func ThirdPublicKey(publicKey string) Option {
	return func(options *Options) {
		options.ThirdPublicKey = publicKey
	}
}

func PrivateKey(privateKey string) Option {
	return func(options *Options) {
		options.PrivateKey = privateKey
	}
}

func Version(version string) Option {
	return func(options *Options) {
		options.Version = version
	}
}