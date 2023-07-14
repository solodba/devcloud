package conf

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 初始化K8S配置
func (k *K8s) GetK8sConfig() (*rest.Config, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", k.KubeConfig)
	if err != nil {
		return nil, err
	}
	return kubeConfig, nil
}

// 获取K8S连接
func (k *K8s) GetK8sConn() (*kubernetes.Clientset, error) {
	k.lock.Lock()
	defer k.lock.Unlock()
	if k.clientSet == nil {
		kubeConfig, err := k.GetK8sConfig()
		if err != nil {
			return nil, err
		}
		clientSet, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return nil, err
		}
		k.clientSet = clientSet
	}
	return k.clientSet, nil
}
