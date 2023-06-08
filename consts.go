package wopan

const (
	ClientID     = "1001000021"
	ClientSecret = "XFmi9GS2hzk98jGX"
	AppID        = "10000001"
	BaseURL      = "https://panservice.mail.wo.cn/"
	DefaultUA    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.37"
)

const (
	APIUserChannel = "api-user"
	WoHomeChannel  = "wohome"
)

// api-user methods
const (
	PcWebLogin        = "PcWebLogin"        // ReqWithParam
	PcLoginVerifyCode = "PcLoginVerifyCode" // ReqWithParam
	AppQueryUser      = "AppQueryUser"      // ReqWithParam
)

// wohome methods
const (
	FCloudProductOrdListQry = "FCloudProductOrdListQry" // ReqWithQry
	QueryCloudUsageInfo     = "QueryCloudUsageInfo"     // ReqWithParamNoClientId
	FCloudProductPackage    = "FCloudProductPackage"    // ReqWithT
	ClassifyRule            = "ClassifyRule"            // ReqWithKey
	GetZoneInfo             = "GetZoneInfo"
	QuerySysConfig          = "QuerySysConfig"
	FamilyUserCurrentEncode = "FamilyUserCurrentEncode"
	QueryAllFiles           = "QueryAllFiles"
)
