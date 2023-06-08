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
	KeyPcWebLogin        = "PcWebLogin"        // ReqWithParam
	KeyPcLoginVerifyCode = "PcLoginVerifyCode" // ReqWithParam
	KeyAppQueryUser      = "AppQueryUser"      // ReqWithParam
)

// wohome methods
const (
	KeyFCloudProductOrdListQry = "FCloudProductOrdListQry" // ReqWithQry
	KeyQueryCloudUsageInfo     = "QueryCloudUsageInfo"     // ReqWithParamNoClientId
	KeyFCloudProductPackage    = "FCloudProductPackage"    // ReqWithT
	KeyClassifyRule            = "ClassifyRule"            // ReqWithKey
	KeyGetZoneInfo             = "GetZoneInfo"
	KeyQuerySysConfig          = "QuerySysConfig"
	KeyFamilyUserCurrentEncode = "FamilyUserCurrentEncode"
	KeyQueryAllFiles           = "QueryAllFiles"
	KeyGetSearchDirectory      = "GetSearchDirectory"
	KeyGetDownloadUrlV2        = "GetDownloadUrlV2"
	KeyGetDownloadUrl          = "GetDownloadUrl"
	KeyCreateDirectory         = "CreateDirectory"
	KeyRenameFileOrDirectory   = "RenameFileOrDirectory"
	KeyMoveFile                = "MoveFile"
	KeyCopyFile                = "CopyFile"
	KeyDeleteFile              = "DeleteFile"
	KeyEmptyRecycleData        = "EmptyRecycleData"
	KeyUpload2C                = "Upload2C"
)

const (
	SortNameAsc = iota + 1
	SortNameDesc
	SortSizeAsc
	SortSizeDesc
	SortTimeAsc
	SortTimeDesc
)
