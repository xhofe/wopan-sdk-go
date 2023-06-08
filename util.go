package wopan

import "reflect"

// CopyStruct copy struct from src to dst by reflect
func CopyStruct(src, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		return
	}
	dstVal = dstVal.Elem()
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		name := srcVal.Type().Field(i).Name
		dstField := dstVal.FieldByName(name)
		// check dst field is valid
		if !dstField.IsValid() {
			continue
		}
		// check dst field is settable
		if !dstField.CanSet() {
			continue
		}
		// check dst field type is same as src field type
		if dstField.Type() != srcField.Type() {
			continue
		}
		dstField.Set(srcField)
	}
}
