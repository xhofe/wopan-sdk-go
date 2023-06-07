package wopan

// FCloudProductOrdListQryData no encrypt
type FCloudProductOrdListQryData struct {
	FcloudProductOrds []struct {
		ActiviteCode        string `json:"activiteCode"`
		AppStorePackageDesc string `json:"appStorePackageDesc"`
		AppStorePackageFee  string `json:"appStorePackageFee"`
		AppStoreProductId   string `json:"appStoreProductId"`
		ApplyTime           string `json:"applyTime"`
		ApplyTimeFormate    string `json:"applyTimeFormate"`
		CbssOrderId         string `json:"cbssOrderId"`
		City                string `json:"city"`
		ClientId            string `json:"clientId"`
		Days                string `json:"days"`
		DescUrl             string `json:"descUrl"`
		EffectState         string `json:"effectState"`
		EffectiveDays       int    `json:"effectiveDays"`
		ExpireTime          string `json:"expireTime"`
		ExpireTimeFormate   string `json:"expireTimeFormate"`
		Fee                 string `json:"fee"`
		IsAppStorePay       string `json:"isAppStorePay"`
		IsAutoSub           string `json:"isAutoSub"`
		IsExpire            string `json:"isExpire"`
		IsNewPackage        string `json:"isNewPackage"`
		IsOnline            string `json:"isOnline"`
		IsPlus              string `json:"isPlus"`
		IsShowExpireTips    string `json:"isShowExpireTips"`
		OrderId             string `json:"orderId"`
		OrderState          string `json:"orderState"`
		OrderStatus         string `json:"orderStatus"`
		PackageDesc         string `json:"packageDesc"`
		PackageProductCode  string `json:"packageProductCode"`
		PackageProductId    string `json:"packageProductId"`
		PayMethod           string `json:"payMethod"`
		PayTransactionId    string `json:"payTransactionId"`
		PayType             string `json:"payType"`
		Province            string `json:"province"`
		RemainDays          string `json:"remainDays"`
		SignStatus          string `json:"signStatus"`
		Source              string `json:"source"`
		SubTime             string `json:"subTime"`
		SubTimeFormate      string `json:"subTimeFormate"`
		SubType             string `json:"subType"`
		UserId              string `json:"userId"`
		VipDesc             string `json:"vipDesc"`
		VipDescNew          string `json:"vipDescNew"`
		VipExpireTimeLabel  string `json:"vipExpireTimeLabel"`
		VipLevel            string `json:"vipLevel"`
	} `json:"fcloudProductOrds"`
	MaxVipLevel string `json:"maxVipLevel"`
	IsShowInlet string `json:"isShowInlet"`
}

type QueryCloudUsageInfoParam struct {
	PhoneNum string `json:"phoneNum"`
	ClientId string `json:"clientId"`
}

type QueryCloudUsageInfoData struct {
	Code      string `json:"code"`
	UsageInfo struct {
		TotalSize     string `json:"totalSize"`
		UsedSize      int    `json:"usedSize"`
		ImageSize     int    `json:"imageSize"`
		VideoSize     int    `json:"videoSize"`
		AudioSize     int    `json:"audioSize"`
		TextSize      int    `json:"textSize"`
		OtherSize     int    `json:"otherSize"`
		ByteUsedSize  int    `json:"byteUsedSize"`
		ByteTotalSize string `json:"byteTotalSize"`
	} `json:"usageInfo"`
	VipLevel   string `json:"vipLevel"`
	ExpireTime string `json:"expireTime"`
	ApplyTime  string `json:"applyTime"`
	PayType    string `json:"payType"`
	Source     string `json:"source"`
	OrderState string `json:"orderState"`
	Status     string `json:"status"`
}

type FCloudProductPackageParam struct {
	ClientId string `json:"clientId"`
	VipLevel string `json:"vipLevel"`
}

type FCloudProductPackageData []struct {
	VipLevel         string `json:"vipLevel"`
	VipDesc          string `json:"vipDesc"`
	FamilyCount      string `json:"familyCount"`
	RecycleBin       string `json:"recycleBin"`
	BatchUploadCount string `json:"batchUploadCount"`
	Capacity         string `json:"capacity"`
	UploadFileSize   string `json:"uploadFileSize"`
	MemberCount      string `json:"memberCount"`
	TransientCount   string `json:"transientCount"`
	VipWeight        string `json:"vipWeight"`
	Packages         []struct {
		AppStoreContractProductId string `json:"appStoreContractProductId"`
		AppStoreDesc              string `json:"appStoreDesc"`
		AppStorePrice             string `json:"appStorePrice"`
		AppStoreProductId         string `json:"appStoreProductId"`
		ApplyTime                 string `json:"applyTime"`
		BatchUploadCount          string `json:"batchUploadCount"`
		Capacity                  string `json:"capacity"`
		Days                      string `json:"days"`
		Desc                      string `json:"desc"`
		DescPic                   string `json:"descPic"`
		ExpireTime                string `json:"expireTime"`
		FamilyCount               string `json:"familyCount"`
		Label                     string `json:"label"`
		MemberCount               string `json:"memberCount"`
		Months                    string `json:"months"`
		Order                     int    `json:"order"`
		OriginalPrice             string `json:"originalPrice"`
		PayType                   string `json:"payType"`
		Price                     string `json:"price"`
		ProductCode               string `json:"productCode"`
		ProductId                 string `json:"productId"`
		RecycleBin                string `json:"recycleBin"`
		Tips                      string `json:"tips"`
		TransientCount            string `json:"transientCount"`
		UploadFileSize            string `json:"uploadFileSize"`
		VipDesc                   string `json:"vipDesc"`
		VipLevel                  string `json:"vipLevel"`
		VipWeight                 string `json:"vipWeight"`
	} `json:"packages"`
}

type ClassifyRuleParam struct {
}
