package utils

import (
	"strconv"
	"strings"
)

func RemoveDuplicatesForInt64(ids []int64) []int64 {
	// 只需要检查key，因此value使用空结构体，不占内存
	processed := make(map[int64]struct{})

	uniqIds := make([]int64, 0)
	for _, uid := range ids {
		// 如果用户ID已经处理过，就跳过
		if _, ok := processed[uid]; ok {
			continue
		}

		// 将唯一的用户ID加到切片中
		uniqIds = append(uniqIds, uid)

		// 将用户ID标记为已存在
		processed[uid] = struct{}{}
	}

	return uniqIds
}

func RemoveDuplicatesForString(uuids []string) []string {
	// 只需要检查key，因此value使用空结构体，不占内存
	processed := make(map[string]struct{})

	uniqUuids := make([]string, 0)
	for _, uuid := range uuids {
		// 如果用户ID已经处理过，就跳过
		if _, ok := processed[uuid]; ok {
			continue
		}

		// 将唯一的用户ID加到切片中
		uniqUuids = append(uniqUuids, uuid)

		// 将用户ID标记为已存在
		processed[uuid] = struct{}{}
	}

	return uniqUuids
}

// SliceCompare 切片比较
func SliceCompareInt64(sliceA, sliceB []int64) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	processedA := make(map[int64]int)
	for _, v := range sliceA {
		if _, ok := processedA[v]; ok {
			processedA[v]++
		} else {
			processedA[v] = 1
		}
	}
	processedB := make(map[int64]int)
	for _, v := range sliceB {
		if _, ok := processedB[v]; ok {
			processedB[v]++
		} else {
			processedB[v] = 1
		}
	}
	for k, va := range processedA {
		if vb, ok := processedB[k]; ok {
			if vb != va {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func SliceCompareString(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	processedA := make(map[string]int)
	for _, v := range sliceA {
		if _, ok := processedA[v]; ok {
			processedA[v]++
		} else {
			processedA[v] = 1
		}
	}
	processedB := make(map[string]int)
	for _, v := range sliceB {
		if _, ok := processedB[v]; ok {
			processedB[v]++
		} else {
			processedB[v] = 1
		}
	}
	for k, va := range processedA {
		if vb, ok := processedB[k]; ok {
			if vb != va {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func SliceFilterInt64(sourceSlice []int64, filterSlice ...int64) []int64 {
	if len(sourceSlice) == 0 {
		return []int64{}
	}
	if len(filterSlice) == 0 {
		return sourceSlice
	}

	processedSource := make(map[int64]int)
	for _, v := range sourceSlice {
		if _, ok := processedSource[v]; ok {
			processedSource[v]++
		} else {
			processedSource[v] = 1
		}
	}
	processedFilter := make(map[int64]int)
	for _, v := range filterSlice {
		if _, ok := processedFilter[v]; ok {
			processedFilter[v]++
		} else {
			processedFilter[v] = 1
		}
	}
	for k, fv := range processedFilter {
		if sv, ok := processedSource[k]; ok {
			n := sv - fv
			if n > 0 {
				processedSource[k] = n
			} else {
				processedSource[k] = 0
			}
		}
	}

	result := make([]int64, 0)
	for k, v := range processedSource {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}

	return result
}

/** SliceFilterString 切片过滤
 */
func SliceFilterString(sourceSlice []string, filterSlice ...string) []string {
	if len(sourceSlice) == 0 {
		return []string{}
	}
	if len(filterSlice) == 0 {
		return sourceSlice
	}

	processedSource := make(map[string]int)
	for _, v := range sourceSlice {
		if _, ok := processedSource[v]; ok {
			processedSource[v]++
		} else {
			processedSource[v] = 1
		}
	}
	processedFilter := make(map[string]int)
	for _, v := range filterSlice {
		if _, ok := processedFilter[v]; ok {
			processedFilter[v]++
		} else {
			processedFilter[v] = 1
		}
	}
	for k, fv := range processedFilter {
		if sv, ok := processedSource[k]; ok {
			n := sv - fv
			if n > 0 {
				processedSource[k] = n
			} else {
				processedSource[k] = 0
			}
		}
	}

	result := make([]string, 0)
	for k, v := range processedSource {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}

	return result
}

func SliceContainInt64(sourceSlice []int64, targetSlice ...int64) bool {
	if len(sourceSlice) < len(targetSlice) {
		return false
	}
	processedSource := make(map[int64]int)
	for _, v := range sourceSlice {
		if _, ok := processedSource[v]; ok {
			processedSource[v]++
		} else {
			processedSource[v] = 1
		}
	}
	processedTarget := make(map[int64]int)
	for _, v := range targetSlice {
		if _, ok := processedTarget[v]; ok {
			processedTarget[v]++
		} else {
			processedTarget[v] = 1
		}
	}
	for k, v := range processedTarget {
		if sv, ok := processedSource[k]; ok {
			if v > sv {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func SliceContainString(sourceSlice []string, targetSlice ...string) bool {
	if len(sourceSlice) < len(targetSlice) {
		return false
	}
	processedSource := make(map[string]int)
	for _, v := range sourceSlice {
		if _, ok := processedSource[v]; ok {
			processedSource[v]++
		} else {
			processedSource[v] = 1
		}
	}
	processedTarget := make(map[string]int)
	for _, v := range targetSlice {
		if _, ok := processedTarget[v]; ok {
			processedTarget[v]++
		} else {
			processedTarget[v] = 1
		}
	}
	for k, v := range processedTarget {
		if sv, ok := processedSource[k]; ok {
			if v > sv {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func SliceIntToString(slice []int64, spacer string) string {
	var strArr []string
	for _, number := range slice {
		s := strconv.FormatInt(number, 10)
		strArr = append(strArr, s)
	}
	return strings.Join(strArr, spacer)
}
