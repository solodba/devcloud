package impl

import (
	"strings"

	"github.com/solodba/devcloud/mkube/apps/pod"
	corev1 "k8s.io/api/core/v1"
)

// k8s中Pod结构体转换成自定义结构体
func (i *impl) PodK8s2Req(podK8s *corev1.Pod) *pod.Pod {
	if podK8s == nil {
		return nil
	}
	return &pod.Pod{
		Base:           i.getReqBase(podK8s),
		Tolerations:    i.getReqTolerations(podK8s.Spec.Tolerations),
		NodeScheduling: i.getNodeReqScheduling(podK8s),
		NetWorking:     i.getReqNetworking(podK8s),
		Volumes:        i.getReqVolumes(podK8s.Spec.Volumes),
		InitContainers: i.getReqContainers(podK8s.Spec.InitContainers),
		Containers:     i.getReqContainers(podK8s.Spec.Containers),
	}
}

// Containers结构体转换
func (i *impl) getReqContainers(containersK8s []corev1.Container) []*pod.Container {
	podReqContainers := make([]*pod.Container, 0)
	for _, item := range containersK8s {
		// container转换
		reqContainer := i.getReqContainer(item)
		podReqContainers = append(podReqContainers, reqContainer)
	}
	return podReqContainers
}

// Container结构体转换
func (i *impl) getReqContainer(container corev1.Container) *pod.Container {
	return &pod.Container{
		Name:            container.Name,
		Image:           container.Image,
		ImagePullPolicy: string(container.ImagePullPolicy),
		Tty:             container.TTY,
		Ports:           i.getReqContainerPorts(container.Ports),
		WorkingDir:      container.WorkingDir,
		Command:         container.Command,
		Args:            container.Args,
		Envs:            i.getReqContainersEnvs(container.Env),
		EnvsFrom:        i.getReqContainersEnvsFrom(container.EnvFrom),
		Privileged:      i.getReqContainerPrivileged(container.SecurityContext),
		Resources:       i.getReqContainerResources(container.Resources),
		VolumeMounts:    i.getReqContainerVolumeMounts(container.VolumeMounts),
		StartupProbe:    i.getReqContainerProbe(container.StartupProbe),
		LivenessProbe:   i.getReqContainerProbe(container.LivenessProbe),
		ReadinessProbe:  i.getReqContainerProbe(container.ReadinessProbe),
	}
}

// ContainerProbe结构体转换
func (i *impl) getReqContainerProbe(probeK8s *corev1.Probe) *pod.ContainerProbe {
	containerProbe := &pod.ContainerProbe{
		Enable: false,
	}
	// 先判断是否探针为空
	if probeK8s != nil {
		// 再判断探针具体是什么类型
		if probeK8s.Exec != nil {
			containerProbe.Enable = true
			containerProbe.Type = PROBE_EXEC
			containerProbe.Exec = &pod.ProbeCommand{
				Command: probeK8s.Exec.Command,
			}
		} else if probeK8s.HTTPGet != nil {
			containerProbe.Enable = true
			containerProbe.Type = PROBE_HTTP
			httpGet := probeK8s.HTTPGet
			headersReq := make([]*pod.ListMapItem, 0)
			for _, headerK8s := range httpGet.HTTPHeaders {
				headersReq = append(headersReq, &pod.ListMapItem{
					Key:   headerK8s.Name,
					Value: headerK8s.Value,
				})
			}
			containerProbe.HttpGet = &pod.ProbeHttpGet{
				Host:        httpGet.Host,
				Port:        httpGet.Port.IntVal,
				Scheme:      string(httpGet.Scheme),
				Path:        httpGet.Path,
				HttpHeaders: headersReq,
			}
		} else if probeK8s.TCPSocket != nil {
			containerProbe.Enable = true
			containerProbe.Type = PROBE_TCP
			containerProbe.TcpSocket = &pod.ProbeTcpSocket{
				Host: probeK8s.TCPSocket.Host,
				Port: probeK8s.TCPSocket.Port.IntVal,
			}
		} else {
			containerProbe.Type = PROBE_HTTP
			return containerProbe
		}
		containerProbe.ProbeTime = &pod.ProbeTime{
			InitialDelaySeconds: probeK8s.InitialDelaySeconds,
			PeriodSeconds:       probeK8s.PeriodSeconds,
			TimeoutSeconds:      probeK8s.TimeoutSeconds,
			SuccessThreshold:    probeK8s.SuccessThreshold,
			FailureThreshold:    probeK8s.FailureThreshold,
		}
	}
	return containerProbe
}

// VolumeMounts结构体转换
func (i *impl) getReqContainerVolumeMounts(volumeMountsK8s []corev1.VolumeMount) []*pod.VolumeMounts {
	volumeMountsReq := make([]*pod.VolumeMounts, 0)
	for _, item := range volumeMountsK8s {
		// 非emptyDir过滤掉
		_, ok := i.volumeMap[item.Name]
		if ok {
			volumeMountsReq = append(volumeMountsReq, &pod.VolumeMounts{
				MountName: item.Name,
				MountPath: item.MountPath,
				ReadOnly:  item.ReadOnly,
			})
		}
	}
	return volumeMountsReq
}

// Resources结构体转换
func (i *impl) getReqContainerResources(requirements corev1.ResourceRequirements) *pod.Resources {
	reqResources := &pod.Resources{
		Enable: false,
	}
	requests := requirements.Requests
	limits := requirements.Limits
	if requests != nil {
		reqResources.Enable = true
		reqResources.CpuRequest = int32(requests.Cpu().MilliValue())
		reqResources.MemRequest = int32(requests.Memory().Value() / (1024 * 1024))
	}
	if limits != nil {
		reqResources.Enable = true
		reqResources.CpuLimit = int32(limits.Cpu().MilliValue())
		reqResources.MemRequest = int32(limits.Memory().Value() / (1024 * 1024))
	}
	return reqResources
}

// Privileged结构体转换
func (i *impl) getReqContainerPrivileged(ctx *corev1.SecurityContext) (privileged bool) {
	if ctx != nil {
		privileged = *ctx.Privileged
	}
	return
}

// Ports结构体转换
func (i *impl) getReqContainerPorts(portsK8s []corev1.ContainerPort) []*pod.ContainerPort {
	portsReq := make([]*pod.ContainerPort, 0)
	for _, item := range portsK8s {
		portsReq = append(portsReq, &pod.ContainerPort{
			Name:          item.Name,
			HostPort:      item.HostPort,
			ContainerPort: item.ContainerPort,
		})
	}
	return portsReq
}

// Envs结构体转换
func (i *impl) getReqContainersEnvs(envsK8s []corev1.EnvVar) []*pod.EnvVar {
	envsReq := make([]*pod.EnvVar, 0)
	for _, item := range envsK8s {
		envVar := &pod.EnvVar{
			Name: item.Name,
		}
		if item.ValueFrom != nil {
			if item.ValueFrom.ConfigMapKeyRef != nil {
				envVar.Type = REF_TYPE_CONFIGMAP
				envVar.Value = item.ValueFrom.ConfigMapKeyRef.Key
				envVar.RefName = item.ValueFrom.ConfigMapKeyRef.Name
				envsReq = append(envsReq, envVar)
			}
			if item.ValueFrom.SecretKeyRef != nil {
				envVar.Type = REF_TYPE_SECRET
				envVar.Value = item.ValueFrom.SecretKeyRef.Key
				envVar.RefName = item.ValueFrom.SecretKeyRef.Name
				envsReq = append(envsReq, envVar)
			}
		} else {
			envVar.Type = "default"
			envVar.Value = item.Value
			envsReq = append(envsReq, envVar)
		}
	}
	return envsReq
}

// EnvsFrom结构体转换
func (i *impl) getReqContainersEnvsFrom(envsFromK8s []corev1.EnvFromSource) []*pod.EnvVarFromResource {
	podReqEnvsFromList := make([]*pod.EnvVarFromResource, 0)
	for _, envK8sItem := range envsFromK8s {
		podReqEnvsFrom := &pod.EnvVarFromResource{
			Prefix: envK8sItem.Prefix,
		}
		if envK8sItem.ConfigMapRef != nil {
			podReqEnvsFrom.RefType = REF_TYPE_CONFIGMAP
			podReqEnvsFrom.Name = envK8sItem.ConfigMapRef.Name
			podReqEnvsFromList = append(podReqEnvsFromList, podReqEnvsFrom)
		}
		if envK8sItem.SecretRef != nil {
			podReqEnvsFrom.RefType = REF_TYPE_SECRET
			podReqEnvsFrom.Name = envK8sItem.SecretRef.Name
			podReqEnvsFromList = append(podReqEnvsFromList, podReqEnvsFrom)
		}
	}
	return podReqEnvsFromList
}

// Volumes结构体转换
func (i *impl) getReqVolumes(volumes []corev1.Volume) []*pod.Volume {
	volumesReq := make([]*pod.Volume, 0)
	for _, volume := range volumes {
		if i.volumeMap == nil {
			i.volumeMap = make(map[string]string)
		}
		var volumeReq *pod.Volume
		if volume.EmptyDir != nil {
			volumeReq = &pod.Volume{
				Type: VOLUME_EMPTYDIR,
				Name: volume.Name,
			}
		}
		if volume.ConfigMap != nil {
			var optional bool
			if volume.ConfigMap.Optional != nil {
				optional = *volume.ConfigMap.Optional
			}
			volumeReq = &pod.Volume{
				Type: VOLUME_CONFIGMAP,
				ConfigMapRefVolume: &pod.ConfigMapRefVolume{
					Name:     volume.ConfigMap.Name,
					Optional: optional,
				},
			}
		}
		if volume.Secret != nil {
			var optional bool
			if volume.Secret.Optional != nil {
				optional = *volume.Secret.Optional
			}
			volumeReq = &pod.Volume{
				Type: VOLUME_SECRET,
				SecretRefVolume: &pod.SecretRefVolume{
					Name:     volume.Secret.SecretName,
					Optional: optional,
				},
			}
		}
		if volume.HostPath != nil {
			var hostPathType corev1.HostPathType
			if volume.HostPath.Type != nil {
				hostPathType = *volume.HostPath.Type
			}
			volumeReq = &pod.Volume{
				Type: VOLUME_HOSTPATH,
				Name: volume.Name,
				HostPathVolume: &pod.HostPathVolume{
					Type: string(hostPathType),
					Path: volume.HostPath.Path,
				},
			}
		}
		if volume.PersistentVolumeClaim != nil {
			volumeReq = &pod.Volume{
				Type: VOLUME_PVC,
				Name: volume.Name,
				PVCVolume: &pod.PVCVolume{
					Name: volume.PersistentVolumeClaim.ClaimName,
				},
			}
		}
		if volume.DownwardAPI != nil {
			downwardAPIVolumeItems := make([]*pod.DownwardAPIVolumeItem, 0)
			for _, item := range volume.DownwardAPI.Items {
				downwardAPIVolumeItems = append(downwardAPIVolumeItems, &pod.DownwardAPIVolumeItem{
					Path:         item.Path,
					FieldRefPath: item.FieldRef.FieldPath,
				})
			}
			volumeReq = &pod.Volume{
				Type: VOLUME_DOWNWARDAPI,
				Name: volume.Name,
				DownwardAPIVolume: &pod.DownwardAPIVolume{
					Items: downwardAPIVolumeItems,
				},
			}
		}
		if volumeReq == nil {
			continue
		}
		i.volumeMap[volume.Name] = ""
		volumesReq = append(volumesReq, volumeReq)
	}
	return volumesReq
}

// NetWorking转换
func (i *impl) getReqNetworking(k8sPod *corev1.Pod) *pod.NetWorking {
	return &pod.NetWorking{
		HostNetwork: k8sPod.Spec.HostNetwork,
		HostName:    k8sPod.Spec.Hostname,
		DnsPolicy:   string(k8sPod.Spec.DNSPolicy),
		DnsConfig:   i.getReqDnsConfig(k8sPod.Spec.DNSConfig),
		HostAliases: i.getReqHostAliases(k8sPod.Spec.HostAliases),
	}
}

// DnsConfig转换
func (i *impl) getReqDnsConfig(dnsConfigK8s *corev1.PodDNSConfig) *pod.DnsConfig {
	dnsConfigReq := &pod.DnsConfig{}
	if dnsConfigK8s != nil {
		dnsConfigReq.Nameservers = dnsConfigK8s.Nameservers
		return dnsConfigReq
	}
	return dnsConfigReq
}

// HostAliases结构体转换
func (i *impl) getReqHostAliases(hostAlias []corev1.HostAlias) []*pod.ListMapItem {
	hostAliasReq := make([]*pod.ListMapItem, 0)
	for _, alias := range hostAlias {
		hostAliasReq = append(hostAliasReq, &pod.ListMapItem{
			Key:   alias.IP,
			Value: strings.Join(alias.Hostnames, ","),
		})
	}
	return hostAliasReq
}

// NodeScheduling结构体转换
func (i *impl) getNodeReqScheduling(podK8s *corev1.Pod) *pod.NodeScheduling {
	nodeScheduling := &pod.NodeScheduling{
		Type: SCHEDULING_NODEANY,
	}
	if podK8s.Spec.NodeSelector != nil {
		nodeScheduling.Type = SCHEDULING_NODESELECTOR
		labels := make([]*pod.ListMapItem, 0)
		for k, v := range podK8s.Spec.NodeSelector {
			labels = append(labels, &pod.ListMapItem{
				Key:   k,
				Value: v,
			})
		}
		nodeScheduling.NodeSelector = labels
		return nodeScheduling
	}
	if podK8s.Spec.Affinity != nil {
		nodeScheduling.Type = SCHEDULING_NODEAFFINITY
		term := podK8s.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms[0]
		matchExpressions := make([]*pod.NodeSelectorTermExpressions, 0)
		for _, expression := range term.MatchExpressions {
			matchExpressions = append(matchExpressions, &pod.NodeSelectorTermExpressions{
				Key:      expression.Key,
				Value:    strings.Join(expression.Values, ""),
				Operator: string(expression.Operator),
			})
		}
		nodeScheduling.NodeAffinity = matchExpressions
		return nodeScheduling
	}
	if podK8s.Spec.NodeName != "" {
		nodeScheduling.Type = SCHEDULING_NODENAME
		nodeScheduling.NodeName = podK8s.Spec.NodeName
		return nodeScheduling
	}
	return nodeScheduling
}

// Tolerations结构体转换
func (i *impl) getReqTolerations(podK8sTolerations []corev1.Toleration) []*pod.Tolerations {
	podReqTolerations := make([]*pod.Tolerations, 0)
	for _, item := range podK8sTolerations {
		podReqTolerations = append(podReqTolerations, &pod.Tolerations{
			Key:      item.Key,
			Value:    item.Value,
			Operator: string(item.Operator),
		})
	}
	return podReqTolerations
}

// Base结构体转换
func (i *impl) getReqBase(podK8s *corev1.Pod) *pod.Base {
	return &pod.Base{
		Name:          podK8s.Name,
		Namespace:     podK8s.Namespace,
		Labels:        i.getReqLabels(podK8s.Labels),
		RestartPolicy: string(podK8s.Spec.RestartPolicy),
	}
}

// Labels转换
func (i *impl) getReqLabels(data map[string]string) []*pod.ListMapItem {
	if data == nil {
		return nil
	}
	labels := make([]*pod.ListMapItem, 0)
	for k, v := range data {
		labels = append(labels, &pod.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return labels
}
