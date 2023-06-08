package wopan

// PcWebLoginData no encrypt
type PcWebLoginData struct {
	NeedSmsCode string `json:"needSmsCode"`
}

func (w *WoClient) PcWebLogin(phone, password string, opts ...RestyOption) (*PcWebLoginData, error) {
	var resp PcWebLoginData
	_, err := w.RequestApiUser("PcWebLogin", Json{
		"phone":        phone,
		"password":     password,
		"uuid":         "",
		"verifyCode":   "",
		"clientSecret": ClientSecret,
	}, JsonClientIDSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PcLoginVerifyCodeData no encrypt
type PcLoginVerifyCodeData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (w *WoClient) PcLoginVerifyCode(phone, password, messageCode string, opts ...RestyOption) (*PcLoginVerifyCodeData, error) {
	var resp PcLoginVerifyCodeData
	_, err := w.RequestApiUser("PcLoginVerifyCode", Json{
		"phone":        phone,
		"messageCode":  messageCode,
		"verifyCode":   nil,
		"uuid":         nil,
		"clientSecret": ClientSecret,
		"password":     password,
	}, JsonClientIDSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AppQueryUserData struct {
	UserId        string `json:"userId"`
	HeadUrl       string `json:"headUrl"`
	UserName      string `json:"userName"`
	Sex           string `json:"sex"`
	Birthday      string `json:"birthday"`
	IsModify      string `json:"isModify"`
	IsHeadModify  string `json:"isHeadModify"`
	IsSetPassword string `json:"isSetPassword"`
	RegisterTime  string `json:"registerTime"`
}

func (w *WoClient) AppQueryUser(opts ...RestyOption) (*AppQueryUserData, error) {
	var resp AppQueryUserData
	_, err := w.RequestApiUser("AppQueryUser", Json{
		"accessToken": w.accessToken,
	}, JsonClientIDSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if w.phone == "" {
		w.phone = resp.UserId
	}
	return &resp, nil
}

// TODO we don't know refresh token method now.
//func (w *WoClient) RefreshToken() ([]byte, error) {
//	res, err := w.RequestApiUser("AppRefreshToken", Json{
//		"refreshToken": w.refreshToken,
//	}, JsonClientIDSecret, nil, ReqOpt{
//		NoAccessToken: true,
//	})
//	return res, err
//}
