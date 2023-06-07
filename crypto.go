package wopan

import "encoding/base64"

type Crypto struct {
	key       []byte
	iv        []byte
	accessKey []byte
}

func NewCrypto() *Crypto {
	c := &Crypto{
		key: []byte("XFmi9GS2hzk98jGX"),
		iv:  []byte("wNSOYIB1k1DjY5lA"),
	}
	return c
}

func (c *Crypto) SetAccessToken(token string) {
	c.accessKey = []byte(token[:16])
}

func (c *Crypto) encrypt(content string, channel string) (string, error) {
	key := c.accessKey
	if channel == "api-user" {
		key = c.key
	}
	res, err := AesEncrypt([]byte(content), key, c.iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

func (c *Crypto) decrypt(content string, channel string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}
	key := c.accessKey
	if channel == "api-user" {
		key = c.key
	}
	res, err := AesDecrypt(bs, key, c.iv)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (c *Crypto) UserEncrypt(content string) (string, error) {
	return c.encrypt(content, "api-user")
}

func (c *Crypto) UserDecrypt(content string) (string, error) {
	return c.decrypt(content, "api-user")
}

func (c *Crypto) Encrypt(content string) (string, error) {
	return c.encrypt(content, "wohome")
}

func (c *Crypto) Decrypt(content string) (string, error) {
	return c.decrypt(content, "wohome")
}
