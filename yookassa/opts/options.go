package opts

import "net/http"

// Config is a mutable options container consumed by yookassa.NewClient.
type Config struct {
	Client *http.Client
}

// Option configures yookassa client creation.
type Option interface {
	Apply(config *Config)
}

type optionFunc func(config *Config)

func (f optionFunc) Apply(config *Config) {
	f(config)
}

// WithHTTPClient overrides default http.Client used by SDK client.
func WithHTTPClient(client http.Client) Option {
	return optionFunc(func(config *Config) {
		copied := client
		config.Client = &copied
	})
}
