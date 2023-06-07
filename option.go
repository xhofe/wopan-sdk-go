package wopan

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Option func(w *WoClient)

func WithAccessToken(token string) Option {
	return func(w *WoClient) {
		w.SetAccessToken(token)
	}
}

func WithRefreshToken(token string) Option {
	return func(w *WoClient) {
		w.SetRefreshToken(token)
	}
}

func WithClient(hc *http.Client) Option {
	return func(c *WoClient) {
		c.SetHttpClient(hc)
	}
}

func WithRestyClient(rc *resty.Client) Option {
	return func(c *WoClient) {
		c.client = rc
	}
}

func WithDebug() Option {
	return func(c *WoClient) {
		c.SetDebug(true)
	}
}

func WithTrace() Option {
	return func(c *WoClient) {
		c.EnableTrace()
	}
}

func WithProxy(proxy string) Option {
	return func(c *WoClient) {
		c.SetProxy(proxy)
	}
}
