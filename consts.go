package wopan

const (
	DefaultClientID     = "1001000021"
	DefaultClientSecret = "XFmi9GS2hzk98jGX"
	DefaultAppID        = "10000001"
	DefaultBaseURL      = "https://panservice.mail.wo.cn"
	DefaultZoneURL      = "https://gxupload.pan.wo.cn:8443"
	DefaultUA           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.37"
	DefaultPartSize     = int64(8 * 1024 * 1024)
)

const (
	ChannelAPIUser = "api-user"
	ChannelWoHome  = "wohome"
	ChannelWoCloud = "wocloud"
)

const (
	SpaceTypePersonal = "0"
	SpaceTypeFamily   = "1"
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
