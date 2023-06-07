package wopan

const (
	ClientID     = "1001000021"
	ClientSecret = "XFmi9GS2hzk98jGX"
	BaseURL      = "https://panservice.mail.wo.cn/"
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
