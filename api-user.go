package wopan

import "context"

// PcWebLoginData no encrypt
type PcWebLoginData struct {
	NeedSmsCode string `json:"needSmsCode"`
}

func (w *WoClient) PcWebLoginCtx(ctx context.Context, phone, password string) (*PcWebLoginData, error) {
	var resp PcWebLoginData
	_, err := w.RequestApiUser("PcWebLogin", Json{
		"phone":        phone,
		"password":     password,
		"uuid":         "",
		"verifyCode":   "",
		"clientSecret": ClientSecret,
	}, JsonClientIDSecret, &resp, ReqOpt{
		NoAccessToken: true,
		Ctx:           ctx,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WoClient) PcWebLogin(phone, password string) (*PcWebLoginData, error) {
	return w.PcWebLoginCtx(context.Background(), phone, password)
}

// PcLoginVerifyCodeData no encrypt
type PcLoginVerifyCodeData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (w *WoClient) PcLoginVerifyCodeCtx(ctx context.Context, phone, password, messageCode string) (*PcLoginVerifyCodeData, error) {
	var resp PcLoginVerifyCodeData
	_, err := w.RequestApiUser("PcLoginVerifyCode", Json{
		"phone":        phone,
		"messageCode":  messageCode,
		"verifyCode":   nil,
		"uuid":         nil,
		"clientSecret": ClientSecret,
		"password":     password,
	}, JsonClientIDSecret, &resp, ReqOpt{
		NoAccessToken: true,
		Ctx:           ctx,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WoClient) PcLoginVerifyCode(phone, password, messageCode string) (*PcLoginVerifyCodeData, error) {
	return w.PcLoginVerifyCodeCtx(context.Background(), phone, password, messageCode)
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

func (w *WoClient) AppQueryUserCtx(ctx context.Context) (*AppQueryUserData, error) {
	var resp AppQueryUserData
	_, err := w.RequestApiUser("AppQueryUser", Json{
		"accessToken": w.accessToken,
	}, JsonClientIDSecret, &resp, ReqOpt{
		Ctx:           ctx,
		NoAccessToken: true,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WoClient) AppQueryUser() (*AppQueryUserData, error) {
	return w.AppQueryUserCtx(context.Background())
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
