/*
这里的方法为slice相关的一些常用方法
但是golang标准库中却没有实现的
*/

package datastructure

import "errors"

var (
	SliceUnsupportedType = errors.New("Unsupported type, Only supported []string, []int, []int64, []int32, []byte. ")
)

//************************************************slice操作********************************************************//

//-------------------------slice去重------------------------------------//
//对于[]string类型去重
func SliceRemoveDupString(slice []string) (reSlice []string, err error) {
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
func SliceRemoveDupInt(slice []int) (reSlice []int, err error) {
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
func SliceRemoveDupInt64(slice []int64) (reSlice []int64, err error) {
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
func SliceRemoveDupInt32(slice []int32) (reSlice []int32, err error) {
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
func SliceRemoveDupByte(slice []byte) (reSlice []byte, err error) {
	m := make(map[byte]struct{})
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			reSlice = append(reSlice, v)
		}
	}
	return
}

//-------------------------slice删除------------------------------------//

var (
	IndexOverflow = errors.New("Index is too small or too lager, this will cause array overflow. ")
)
//根据index删除
func SliceDeleteStringEleByIndex(slice []string, index int) (reSlice []string, err error) {
	if index < 0 || index >= len(slice) {
		err = IndexOverflow
		return
	}
	reSlice = append(slice[:index], slice[index + 1:]...)
	return
}

func SliceDeleteIntEleByIndex(slice []int, index int) (reSlice []int, err error) {
	if index < 0 || index >= len(slice) {
		err = IndexOverflow
		return
	}
	reSlice = append(slice[:index], slice[index + 1:]...)
	return
}

func SliceDeleteInt64EleByIndex(slice []int64, index int) (reSlice []int64, err error) {
	if index < 0 || index >= len(slice) {
		err = IndexOverflow
		return
	}
	reSlice = append(slice[:index], slice[index + 1:]...)
	return
}

func SliceDeleteInt32EleByIndex(slice []int32, index int) (reSlice []int32, err error) {
	if index < 0 || index >= len(slice) {
		err = IndexOverflow
		return
	}
	reSlice = append(slice[:index], slice[index + 1:]...)
	return
}

func SliceDeleteByteEleByIndex(slice []byte, index int) (reSlice []byte, err error) {
	if index < 0 || index >= len(slice) {
		err = IndexOverflow
		return
	}
	reSlice = append(slice[:index], slice[index + 1:]...)
	return
}

//-------------------------slice判断是否相等------------------------------------//

