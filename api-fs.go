package wopan

type QueryAllFilesData struct {
	Files []struct {
		FamilyId     int    `json:"familyId"`
		Fid          string `json:"fid"`
		Creator      string `json:"creator"`
		Size         int    `json:"size"`
		CreateTime   string `json:"createTime"`
		Name         string `json:"name"`
		ShootingTime string `json:"shootingTime"`
		Id           string `json:"id"`
		Type         int    `json:"type"`
		ThumbUrl     string `json:"thumbUrl"`
		FileType     string `json:"fileType"`
	} `json:"files"`
}

func (w *WoClient) QueryAllFiles(spaceType, parentDirectoryId string, pageNum, pageSize int, sortRule int, familyId string, opts ...RestyOption) (*QueryAllFilesData, error) {
	var resp QueryAllFilesData
	param := Json{
		"spaceType":         spaceType,
		"parentDirectoryId": parentDirectoryId,
		"pageNum":           pageNum,
		"pageSize":          pageSize,
		"sortRule":          sortRule,
		"clientId":          ClientID,
	}
	if spaceType == "1" {
		param["familyId"] = familyId
	}
	_, err := w.RequestWoHome(KeyQueryAllFiles, param, JsonSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WoClient) QueryAllFilesPersonal(parentDirectoryId string, pageNum, pageSize int, sortRule int, opts ...RestyOption) (*QueryAllFilesData, error) {
	return w.QueryAllFiles("0", parentDirectoryId, pageNum, pageSize, sortRule, "", opts...)
}

func (w *WoClient) QueryAllFilesFamily(parentDirectoryId string, pageNum, pageSize int, sortRule int, familyId string, opts ...RestyOption) (*QueryAllFilesData, error) {
	return w.QueryAllFiles("1", parentDirectoryId, pageNum, pageSize, sortRule, familyId, opts...)
}

// GetSearchDirectory??

type GetDownloadUrlV2Data struct {
	Type int `json:"type"`
	List []struct {
		Fid         string `json:"fid"`
		DownloadUrl string `json:"downloadUrl"`
	} `json:"list"`
}

func (w *WoClient) GetDownloadUrlV2(fidList []string, opts ...RestyOption) (*GetDownloadUrlV2Data, error) {
	var resp GetDownloadUrlV2Data
	param := Json{
		"type":     "1",
		"fidList":  fidList,
		"clientId": ClientID,
	}
	_, err := w.RequestWoHome(KeyGetDownloadUrlV2, param, JsonSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDownloadUrlData struct {
	Fid         string `json:"fid"`
	DownloadUrl string `json:"downloadUrl"`
}

func (w *WoClient) GetDownloadUrl(spaceType string, fidList []string, opts ...RestyOption) (*GetDownloadUrlData, error) {
	var resp GetDownloadUrlData
	param := Json{
		"fidList":   fidList,
		"clientId":  "1001000001", // ???
		"spaceType": spaceType,
	}
	_, err := w.RequestWoHome(KeyGetDownloadUrl, param, JsonSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDirectoryData struct {
	Id string `json:"id"`
}

func (w *WoClient) CreateDirectory(spaceType, parentDirectoryId string, directoryName, familyId string, opts ...RestyOption) (*CreateDirectoryData, error) {
	var resp CreateDirectoryData
	param := Json{
		"spaceType":         spaceType,
		"familyId":          familyId,
		"parentDirectoryId": parentDirectoryId,
		"directoryName":     directoryName,
		"clientId":          ClientID,
	}
	_, err := w.RequestWoHome(KeyCreateDirectory, param, JsonSecret, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
