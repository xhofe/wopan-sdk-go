package wopan

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type WoClient struct {
	client       *resty.Client
	crypto       *Crypto
	accessToken  string
	refreshToken string
}

func New(opts ...Option) *WoClient {
	w := &WoClient{
		client: resty.New(),
		crypto: NewCrypto(),
	}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

func Default(refreshToken string) *WoClient {
	return New(WithRefreshToken(refreshToken))
}

func (w *WoClient) SetAccessToken(token string) {
	w.accessToken = token
	w.crypto.SetAccessToken(token)
}

func (w *WoClient) SetRefreshToken(token string) {
	w.refreshToken = token
}

func (w *WoClient) SetHttpClient(httpClient *http.Client) *WoClient {
	w.client = resty.NewWithClient(httpClient)
	return w
}

func (w *WoClient) SetUserAgent(userAgent string) *WoClient {
	w.client.SetHeader("User-Agent", userAgent)
	return w
}

func (w *WoClient) SetCookies(cs ...*http.Cookie) *WoClient {
	w.client.SetCookies(cs)
	return w
}

func (w *WoClient) SetDebug(d bool) *WoClient {
	w.client.SetDebug(d)
	return w
}

func (w *WoClient) EnableTrace() *WoClient {
	w.client.EnableTrace()
	return w
}

func (w *WoClient) SetProxy(proxy string) *WoClient {
	w.client.SetProxy(proxy)
	return w
}

func (w *WoClient) NewRequest() *resty.Request {
	return w.client.R()
}
