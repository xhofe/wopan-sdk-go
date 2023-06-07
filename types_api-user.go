package wopan

type PcWebLoginParam struct {
	Phone        string      `json:"phone"`
	Password     string      `json:"password"`
	Uuid         interface{} `json:"uuid"`
	VerifyCode   interface{} `json:"verifyCode"`
	ClientSecret string      `json:"clientSecret"`
}

// PcWebLoginData no encrypt
type PcWebLoginData struct {
	NeedSmsCode string `json:"needSmsCode"`
}

type PcLoginVerifyCodeParam struct {
	PcWebLoginParam
	MessageCode string `json:"messageCode"`
}

// PcLoginVerifyCodeData no encrypt
type PcLoginVerifyCodeData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type AppQueryUserParam struct {
	AccessToken string `json:"accessToken"`
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
