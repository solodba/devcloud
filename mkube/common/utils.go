package common

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"time"

	"github.com/solodba/devcloud/mkube/apps/rbac"
)

// 机构体转换成映射
func ToMap(items []*rbac.ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range items {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}

// 映射转换成结构体
func ToList(data map[string]string) []*rbac.ListMapItem {
	list := make([]*rbac.ListMapItem, 0)
	for k, v := range data {
		list = append(list, &rbac.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return list
}

// 映射转换成结构体
func ToListWithMapByte(data map[string][]byte) []*rbac.ListMapItem {
	list := make([]*rbac.ListMapItem, 0)
	for k, v := range data {
		list = append(list, &rbac.ListMapItem{
			Key:   k,
			Value: string(v),
		})
	}
	return list
}

// 基于字符串的哈希码生成 RGB 字符串
func GenerateHashBasedRGB(str string) string {
	hash := HashString(str)    // 计算字符串的哈希码
	r, g, b := HashToRGB(hash) // 将哈希码转换为 RGB 分量

	return strconv.Itoa(r) + "," + strconv.Itoa(g) + "," + strconv.Itoa(b)
}

// 计算字符串的哈希码
func HashString(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

// 将哈希码转换为 RGB 分量
func HashToRGB(hash uint32) (r, g, b int) {
	r = int(hash & 0xFF)         // 取低 8 位作为红色分量
	g = int((hash >> 8) & 0xFF)  // 取中间 8 位作为绿色分量
	b = int((hash >> 16) & 0xFF) // 取高 8 位作为蓝色分量
	return
}

func GetFormatTimeByUnix(createTime int64) string {
	if createTime == 0 {
		return "Unknown"
	}
	//计算时间
	currentTime := time.Now()
	timestampTime := time.Unix(createTime, 0)
	days := int(currentTime.Sub(timestampTime).Hours() / 24)

	years := days / 365         // 计算年份
	remainingDays := days % 365 // 剩余的天数

	months := remainingDays / 30       // 计算月份
	remainingDays = remainingDays % 30 // 剩余的天数

	result := ""
	if years > 0 {
		result += fmt.Sprintf("%d年", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%d月", months)
	}
	if remainingDays > 0 {
		result += fmt.Sprintf("%d天", remainingDays)
	}
	return result
}
