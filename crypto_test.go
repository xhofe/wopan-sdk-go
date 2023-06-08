package wopan_test

import (
	"testing"

	"github.com/Xhofe/wopan-sdk-go"
)

func TestCrypto_Decrypt(t *testing.T) {
	c := wopan.NewCrypto()
	c.SetAccessToken("c3f0362b-0515-4997-a9c2-0478702a5fea")
	res, err := c.WoHomeDecrypt("b+fh7Uuv5WqECmt1B++feofiJE8QlaGAN0GEFe9UYlLgLetBsq0Vq3bfPJQlT89h/Nzfye/UwN/gOSEMS/GYUw0NVilHlz+gFRW5x0rtpsdqjcIF7B6qQnE8mYy2oLrwxku5XHZ3lwBsMdMg6yrwhviXg4I1cqKEAKE94xxMDUaOHAGxSwHmleOLJA9TDt7h5PRA6USeRQrdy/HIambvwY56UK8nwgdocAC/H19DR0hlPko7DAfTpSDps9rEBDebvlk+G88WF5f7S/LgOOwsULeSEf9WGqj+kO+shOkhj5T5IR+AQ/rQqCT8WadKgbig")
	if err != nil {
		t.Errorf("WoHomeDecrypt() error = %v", err)
	} else {
		t.Logf("WoHomeDecrypt() res = %v", res)
	}
}

func TestCrypto_UserDecrypt(t *testing.T) {
	c := wopan.NewCrypto()
	res, err := c.UserDecrypt("OUiQMdAaILKTyrcyS0RSd9Kdd86NnTndTnZ45C1iz1qqLlaMbPA7IEhEsyowFNSgj5lsTwXTWnA8e/gZsV1QW39oK1ZUn0Nw8ACNwHgUNsXSc/UjZ2vPf+mlXqCN2HL9JQLEddvvZER319qeaxrHI1A/t7XZOgX9YChhO9QR1tiakwLxKWaIDH4AnAmR7WSBmo2dWFoRzvJBibTCD3eSAjqGX6hOsZ2h5qyINDtS303qWhQRsYPIu3AVQ9RSub5OJpiLPUSIKS1vNQ9r/m2AgT5lLoSfNkOdoIHNhWbJqfC1elT1JC4IlybttsXQPqj60FK36ujvSLwsvXse+SKa2Q==")
	if err != nil {
		t.Errorf("UserDecrypt() error = %v", err)
	} else {
		t.Logf("UserDecrypt() res = %v", res)
	}
}

func TestCrypto_UserEncrypt(t *testing.T) {
	c := wopan.NewCrypto()
	res, err := c.UserEncrypt(`{"userId":"186****5244","headUrl":"https://panservice.mail.wo.cn/upload/headPortrait/person_1600762418973144.png","userName":"wo_Kbc7Hb","sex":"","birthday":"","isModify":"0","isHeadModify":"0","isSetPassword":"1","registerTime":"20220221185416"}`)
	if err != nil {
		t.Errorf("UserEncrypt() error = %v", err)
	} else {
		t.Logf("UserEncrypt() res = %v", res)
	}
}
