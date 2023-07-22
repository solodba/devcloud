package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/node"
	corev1 "k8s.io/api/core/v1"
)

// 获取k8s节点状态
func (i *impl) getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	nodeStatus := "NotReady"
	for _, condition := range nodeConditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			nodeStatus = "Ready"
			break
		}
	}
	return nodeStatus
}

// 获取k8s节点信息中External-IP和Internal-IP
func (i *impl) getNodeIp(addresses []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
	for _, item := range addresses {
		if item.Type == addressType {
			return item.Address
		}
	}
	return "<node>"
}

// 把k8s节点信息结构体转换成response的Node结构体
func (i *impl) GetNodeResItem(nodeK8s corev1.Node) *node.NodeSetItem {
	nodeInfo := nodeK8s.Status.NodeInfo
	return &node.NodeSetItem{
		Name:             nodeK8s.Name,
		Status:           i.getNodeStatus(nodeK8s.Status.Conditions),
		Age:              nodeK8s.CreationTimestamp.Unix(),
		InternalIp:       i.getNodeIp(nodeK8s.Status.Addresses, corev1.NodeInternalIP),
		ExternalIp:       i.getNodeIp(nodeK8s.Status.Addresses, corev1.NodeExternalIP),
		OsImage:          nodeInfo.OSImage,
		Version:          nodeInfo.KubeletVersion,
		KernelVersion:    nodeInfo.KernelVersion,
		ContainerRuntime: nodeInfo.ContainerRuntimeVersion,
	}
}

// 获取Node详情
func (i *impl) GetNodeDetail(nodeK8s corev1.Node) *node.NodeSetItem {
	nodeRes := i.GetNodeResItem(nodeK8s)
	// 计算label和taint
	nodeRes.Taints = i.getTaintReq(nodeK8s.Spec.Taints)
	nodeRes.Labels = i.mapToList(nodeK8s.Labels)
	return nodeRes
}

// Taint转换
func (i *impl) getTaintReq(k8sTaint []corev1.Taint) []*node.Taint {
	taintReqList := make([]*node.Taint, 0)
	for _, item := range k8sTaint {
		taintReqList = append(taintReqList, &node.Taint{
			Key:    item.Key,
			Value:  item.Value,
			Effect: string(item.Effect),
		})
	}
	return taintReqList
}

// Labels转换
func (i *impl) mapToList(m map[string]string) []*node.ListMapItem {
	res := make([]*node.ListMapItem, 0)
	for k, v := range m {
		res = append(res, &node.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return res
}
