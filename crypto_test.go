package wopan_test

import (
	"testing"

	"github.com/Xhofe/wopan-sdk-go"
)

func TestCrypto_Decrypt(t *testing.T) {
	c := wopan.NewCrypto()
	c.SetAccessToken("91d4b946-xxxx-4909-bac1-d9914e45f2de")
	res, err := c.WoHomeDecrypt("JcACIY9uPch9xcdwNMFYBGrr6ZS+Fh5bcKe2I8IxjTZb1hDGWya9nWpaX2obH63RskoDtGBeEswRrVX5ZQFxLKwHNE4AHbkXtfmilh6MXe05Jzcrv/cMqCeCt05GjyolqoyzzZtrLTS/FT8rvftMTmuoPNzKBYNr8uV5ZytedJDVK3dYOnaL1n+efCVQxDQv25uMEgM4JP4rAy/v+4IUDGO+zGyjj3YvmLT4Y99aTo2ue2acOX9NrfjsHs1dDq7Q")
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
