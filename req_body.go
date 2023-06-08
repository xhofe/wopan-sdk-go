package wopan

func (w *WoClient) NewBody(channel string, param, other Json) (Json, error) {
	if param == nil {
		return other, nil
	}
	jsonBytes, err := w.jsonMarshalFunc(param)
	if err != nil {
		return nil, err
	}
	encrypted, err := w.crypto.EncryptBytes(jsonBytes, channel)
	if err != nil {
		return nil, err
	}
	other["param"] = encrypted
	return other, nil
}
