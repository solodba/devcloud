package common

import "github.com/solodba/devcloud/mpaas/apps/rbac"

func ToMap(items []*rbac.ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range items {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}
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
