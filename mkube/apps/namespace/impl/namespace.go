package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mkube/apps/namespace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 查询Namespace集合
func (i *impl) QueryNamespace(ctx context.Context, in *namespace.QueryNamespaceRequest) (*namespace.NamespaceSet, error) {
	namespaceApi := i.clientSet.CoreV1().Namespaces()
	namespaceList, err := namespaceApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get namespace list error, err: %s", err.Error())
	}
	namespaceSet := namespace.NewNamespaceSet()
	for _, item := range namespaceList.Items {
		namespace := i.NamespaceK8s2ResItemConvert(item)
		namespaceSet.AddItems(namespace)
	}
	namespaceSet.Total = int64(len(namespaceSet.Items))
	return namespaceSet, nil
}
