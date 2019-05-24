/*
这里的方法为slice相关的一些常用方法
但是golang标准库中却没有实现的
*/

package datastructure

import "errors"

var (
	SliceUnsupportedType = errors.New("Unsupported type, Only supported []string, []int, []int64, []int32, []byte")
)


//slice去重方法
func SliceRemoveDup(slice interface{}) (reSilce []interface{}, err error) {
	switch slice.(type) {
	case []string:
	case []int:
	case []int64:
	case []int32:
	case []byte:
	default:
		return nil, SliceUnsupportedType
	}
	return nil, nil
}

//slice类型转换

//对于[]string类型去重
func RemoveDupString(slice []string) (reSlice []string, err error) {
	m := make(map[string]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}

//对于[]int类型去重
func RemoveDupInt(slice []int) (reSlice []int, err error) {
	m := make(map[int]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}

//对于[]int64类型去重
func RemoveDupInt64(slice []int64) (reSlice []int64, err error) {
	m := make(map[int64]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}

//对于[]int32类型去重
func RemoveDupInt32(slice []int32) (reSlice []int32, err error) {
	m := make(map[int32]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}

//对于[]byte类型去重
func RemoveDupByte(slice []byte) (reSlice []byte, err error) {
	m := make(map[byte]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}
