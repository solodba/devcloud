package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 查询集群资源使用情况
func (i *impl) QueryClusterUsage(ctx context.Context, in *metrics.QueryClusterUsageRequest) (*metrics.MetricSet, error) {
	metricSet := metrics.NewMetricSet()
	raw, err := i.clientSet.RESTClient().Get().AbsPath(QUERY_CLUSTER_USAGE).DoRaw(ctx)
	if err != nil {
		return metricSet, err
	}
	nodeMetricSet := metrics.NewNodeMetricSet()
	err = json.Unmarshal(raw, nodeMetricSet)
	if err != nil {
		return metricSet, err
	}
	nodeList, err := i.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return metricSet, err
	}
	if len(nodeList.Items) != len(nodeMetricSet.Items) {
		return metricSet, err
	}
	var cpuUsage, cpuTotal int64
	var memUsage, memTotal int64
	podList, err := i.clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return metricSet, err
	}
	var podUsage, podTotal int64 = int64(len(podList.Items)), 0
	for i, item := range nodeList.Items {
		cpuUsage += nodeMetricSet.Items[i].Usage.Cpu().MilliValue()
		cpuTotal += item.Status.Capacity.Cpu().MilliValue()
		memUsage += nodeMetricSet.Items[i].Usage.Memory().Value()
		memTotal += item.Status.Capacity.Memory().Value()
		podTotal += item.Status.Capacity.Pods().Value()
	}
	podUsageFormat := fmt.Sprintf("%.2f", (float64(podUsage)/float64(podTotal))*100)
	podMetricItem := metrics.NewMetricItem()
	podMetricItem.Value = podUsageFormat
	podMetricItem.Title = "Pod使用占比"
	cpuUsageFormat := fmt.Sprintf("%.2f", (float64(cpuUsage)/float64(cpuTotal))*100)
	cpuMetricItem := metrics.NewMetricItem()
	cpuMetricItem.Value = cpuUsageFormat
	cpuMetricItem.Title = "CPU使用占比"
	cpuMetricItem.Label = "cluster_cpu"
	memUsageFormat := fmt.Sprintf("%.2f", (float64(memUsage)/float64(memTotal))*100)
	memMetricItem := metrics.NewMetricItem()
	memMetricItem.Value = memUsageFormat
	memMetricItem.Title = "Memory使用占比"
	memMetricItem.Label = "cluster_mem"
	metricSet.AddItems(podMetricItem, cpuMetricItem, memMetricItem)
	metricSet.Total = int64(len(metricSet.Items))
	// 写入数据库
	_, err = i.col.InsertOne(ctx, metricSet)
	if err != nil {
		return nil, err
	}
	return metricSet, nil
}

// 查询K8S资源情况
func (i *impl) QueryResource(ctx context.Context, in *metrics.QueryResourceRequest) (*metrics.MetricSet, error) {
	metricSet := metrics.NewMetricSet()
	// 获取namespace
	namespaceList, err := i.clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err == nil {
		namespaceMetric := metrics.NewMetricItem()
		namespaceMetric.Logo = "k8s"
		namespaceMetric.Title = "Namespaces"
		namespaceMetric.Value = strconv.Itoa(len(namespaceList.Items))
		metricSet.AddItems(namespaceMetric)
	}
	// 获取Pods
	podList, err := i.clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err == nil {
		podMetric := metrics.NewMetricItem()
		podMetric.Logo = "pod"
		podMetric.Title = "Pods"
		podMetric.Value = strconv.Itoa(len(podList.Items))
		metricSet.AddItems(podMetric)
	}
	// 获取ConfigMap
	cmList, err := i.clientSet.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err == nil {
		cmMetric := metrics.NewMetricItem()
		cmMetric.Logo = "cm"
		cmMetric.Title = "ConfigMaps"
		cmMetric.Value = strconv.Itoa(len(cmList.Items))
		metricSet.AddItems(cmMetric)
	}
	// 获取Secret
	secretList, err := i.clientSet.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		secretMetric := metrics.NewMetricItem()
		secretMetric.Logo = "secret"
		secretMetric.Title = "Secrets"
		secretMetric.Value = strconv.Itoa(len(secretList.Items))
		metricSet.AddItems(secretMetric)
	}
	// 获取PersistentVolume
	pvList, err := i.clientSet.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err == nil {
		pvMetric := metrics.NewMetricItem()
		pvMetric.Logo = "pv"
		pvMetric.Title = "PV"
		pvMetric.Value = strconv.Itoa(len(pvList.Items))
		metricSet.AddItems(pvMetric)
	}
	// 获取PersistentVolumeClaim
	pvcList, err := i.clientSet.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err == nil {
		pvcMetric := metrics.NewMetricItem()
		pvcMetric.Logo = "pvc"
		pvcMetric.Title = "PVC"
		pvcMetric.Value = strconv.Itoa(len(pvcList.Items))
		metricSet.AddItems(pvcMetric)
	}
	// 获取StorageClass
	scList, err := i.clientSet.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err == nil {
		scMetric := metrics.NewMetricItem()
		scMetric.Logo = "sc"
		scMetric.Title = "StorageClass"
		scMetric.Value = strconv.Itoa(len(scList.Items))
		metricSet.AddItems(scMetric)
	}
	// 获取Service
	svcList, err := i.clientSet.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err == nil {
		svcMetric := metrics.NewMetricItem()
		svcMetric.Logo = "svc"
		svcMetric.Title = "Services"
		svcMetric.Value = strconv.Itoa(len(svcList.Items))
		metricSet.AddItems(svcMetric)
	}
	// 获取Ingress
	ingressList, err := i.clientSet.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err == nil {
		ingressMetric := metrics.NewMetricItem()
		ingressMetric.Logo = "ingress"
		ingressMetric.Title = "Ingresses"
		ingressMetric.Value = strconv.Itoa(len(ingressList.Items))
		metricSet.AddItems(ingressMetric)
	}
	// 获取Deployment
	deploymentList, err := i.clientSet.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err == nil {
		deploymentMetric := metrics.NewMetricItem()
		deploymentMetric.Logo = "pod"
		deploymentMetric.Title = "Deployments"
		deploymentMetric.Value = strconv.Itoa(len(deploymentList.Items))
		metricSet.AddItems(deploymentMetric)
	}
	// 获取DaemonSet
	daemonSetList, err := i.clientSet.AppsV1().DaemonSets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		daemonSetMetric := metrics.NewMetricItem()
		daemonSetMetric.Logo = "pod"
		daemonSetMetric.Title = "DaemonSets"
		daemonSetMetric.Value = strconv.Itoa(len(daemonSetList.Items))
		metricSet.AddItems(daemonSetMetric)
	}
	// 获取StatefulSets
	statefulSetList, err := i.clientSet.AppsV1().StatefulSets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		statefulSetMetric := metrics.NewMetricItem()
		statefulSetMetric.Logo = "ingress"
		statefulSetMetric.Title = "statefulSets"
		statefulSetMetric.Value = strconv.Itoa(len(statefulSetList.Items))
		metricSet.AddItems(statefulSetMetric)
	}
	// 获取Jobs
	jobList, err := i.clientSet.BatchV1().Jobs("").List(ctx, metav1.ListOptions{})
	if err == nil {
		jobMetric := metrics.NewMetricItem()
		jobMetric.Logo = "pod"
		jobMetric.Title = "Jobs"
		jobMetric.Value = strconv.Itoa(len(jobList.Items))
		metricSet.AddItems(jobMetric)
	}
	// 获取CronJobs
	cronJobList, err := i.clientSet.BatchV1().CronJobs("").List(ctx, metav1.ListOptions{})
	if err == nil {
		cronjobMetric := metrics.NewMetricItem()
		cronjobMetric.Logo = "pod"
		cronjobMetric.Title = "CronJobs"
		cronjobMetric.Value = strconv.Itoa(len(cronJobList.Items))
		metricSet.AddItems(cronjobMetric)
	}
	// 获取ServiceAccounts
	saList, err := i.clientSet.CoreV1().ServiceAccounts("").List(ctx, metav1.ListOptions{})
	if err == nil {
		saMetric := metrics.NewMetricItem()
		saMetric.Logo = "secret"
		saMetric.Title = "ServiceAccounts"
		saMetric.Value = strconv.Itoa(len(saList.Items))
		metricSet.AddItems(saMetric)
	}
	// 获取Roles
	roleList, err := i.clientSet.RbacV1().Roles("").List(ctx, metav1.ListOptions{})
	if err == nil {
		roleMetric := metrics.NewMetricItem()
		roleMetric.Logo = "secret"
		roleMetric.Title = "Roles"
		roleMetric.Value = strconv.Itoa(len(roleList.Items))
		metricSet.AddItems(roleMetric)
	}
	// 获取ClusterRole
	clusterRoleList, err := i.clientSet.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err == nil {
		clusterRoleMetric := metrics.NewMetricItem()
		clusterRoleMetric.Logo = "secret"
		clusterRoleMetric.Title = "ClusterRoles"
		clusterRoleMetric.Value = strconv.Itoa(len(clusterRoleList.Items))
		metricSet.AddItems(clusterRoleMetric)
	}
	// 获取RoleBinding
	roleBindingList, err := i.clientSet.RbacV1().RoleBindings("").List(ctx, metav1.ListOptions{})
	if err == nil {
		roleBindingMetric := metrics.NewMetricItem()
		roleBindingMetric.Logo = "secret"
		roleBindingMetric.Title = "RoleBindings"
		roleBindingMetric.Value = strconv.Itoa(len(roleBindingList.Items))
		metricSet.AddItems(roleBindingMetric)
	}
	// 获取ClusterRoleBinding
	clusterRoleBindingList, err := i.clientSet.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err == nil {
		clusterRoleBindingMetric := metrics.NewMetricItem()
		clusterRoleBindingMetric.Logo = "secret"
		clusterRoleBindingMetric.Title = "CRBindings"
		clusterRoleBindingMetric.Value = strconv.Itoa(len(clusterRoleBindingList.Items))
		metricSet.AddItems(clusterRoleBindingMetric)
	}
	// 随机Color的生成
	for index, item := range metricSet.Items {
		metricSet.Items[index].Color = common.GenerateHashBasedRGB(item.Value)
	}
	return metricSet, nil
}

// 获取集群信息
func (i *impl) QueryClusterInfo(ctx context.Context, in *metrics.QueryClusterInfoRequest) (*metrics.MetricSet, error) {
	metricSet := metrics.NewMetricSet()
	// 添加K8S类型信息
	typeMetric := metrics.NewMetricItem()
	typeMetric.Title = "Cluster"
	typeMetric.Logo = "k8s"
	typeMetric.Value = "K8S"
	metricSet.AddItems(typeMetric)
	// 添加K8S版本信息
	serverVersion, err := i.clientSet.ServerVersion()
	if err == nil {
		versionMetric := metrics.NewMetricItem()
		versionMetric.Title = "Kubernetes version"
		versionMetric.Logo = "k8s"
		versionMetric.Value = fmt.Sprintf("%s.%s", serverVersion.Major, serverVersion.Minor)
		metricSet.AddItems(versionMetric)
	}
	// K8S集群初始化时间
	nodeList, err := i.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err == nil {
		var k8sCreateTime int64 = 0
		for _, item := range nodeList.Items {
			if _, ok := item.Labels["node-role.kubernetes.io/control-plane"]; ok {
				if k8sCreateTime == 0 {
					k8sCreateTime = item.CreationTimestamp.Unix()
				}
				if k8sCreateTime > 0 && item.CreationTimestamp.Unix() < k8sCreateTime {
					k8sCreateTime = item.CreationTimestamp.Unix()
				}
			}
		}
		createTimeMetric := metrics.NewMetricItem()
		createTimeMetric.Title = "Created"
		createTimeMetric.Logo = "k8s"
		createTimeMetric.Value = common.GetFormatTimeByUnix(k8sCreateTime)
		metricSet.AddItems(createTimeMetric)
	}
	// K8S节点数量
	if err == nil {
		nodeCountMetric := metrics.NewMetricItem()
		nodeCountMetric.Title = "Nodes"
		nodeCountMetric.Logo = "k8s"
		nodeCountMetric.Value = strconv.Itoa(len(nodeList.Items))
		metricSet.AddItems(nodeCountMetric)
	}
	// 随机Color的生成
	for index, item := range metricSet.Items {
		metricSet.Items[index].Color = common.GenerateHashBasedRGB(item.Title)
	}
	return metricSet, nil
}
