package impl

import (
	"strconv"
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/pod"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// 定义常量
const (
	SCHEDULING_NODENAME     = "nodeName"
	SCHEDULING_NODESELECTOR = "nodeSelector"
	SCHEDULING_NODEAFFINITY = "nodeAffinity"
	SCHEDULING_NODEANY      = "nodeAny"
)

const (
	REF_TYPE_CONFIGMAP = "configMap"
	REF_TYPE_SECRET    = "secret"
)

const (
	PROBE_HTTP = "http"
	PROBE_TCP  = "tcp"
	PROBE_EXEC = "exec"
)

const (
	VOLUME_EMPTYDIR    = "emptyDir"
	VOLUME_CONFIGMAP   = "configMap"
	VOLUME_SECRET      = "secret"
	VOLUME_HOSTPATH    = "hostPath"
	VOLUME_DOWNWARDAPI = "downwardAPI"
	VOLUME_PVC         = "pvc"
)

// Pod自定义结构体转换成K8s中定义的结构体
func (i *impl) PodReq2K8s(p *pod.Pod) *corev1.Pod {
	nodeAffinity, nodeSelector, nodeName := i.getNodeK8sScheduling(p)
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Base.Name,
			Namespace: p.Base.Namespace,
			Labels:    i.getK8sLabels(p),
		},
		Spec: corev1.PodSpec{
			NodeName:       nodeName,
			NodeSelector:   nodeSelector,
			Affinity:       nodeAffinity,
			Tolerations:    i.getK8sTolerations(p),
			InitContainers: i.getK8sContainer(p.InitContainers),
			Containers:     i.getK8sContainer(p.Containers),
			Volumes:        i.getK8sVolumes(p),
			DNSConfig:      i.getK8sDnsConfig(p),
			DNSPolicy:      i.getK8sDNSPolicy(p),
			HostAliases:    i.getK8sHostAlias(p),
			Hostname:       i.getK8sHostname(p),
			RestartPolicy:  corev1.RestartPolicy(p.Base.RestartPolicy),
		},
	}
}

// Hostname转换
func (i *impl) getK8sHostname(p *pod.Pod) string {
	var k8sHostname string
	if p.NetWorking == nil {
		return k8sHostname
	}
	k8sHostname = p.NetWorking.HostName
	return k8sHostname
}

// DNSConfig转换
func (i *impl) getK8sDnsConfig(p *pod.Pod) *corev1.PodDNSConfig {
	if p.NetWorking == nil {
		return nil
	}
	if p.NetWorking.DnsConfig == nil {
		return nil
	}
	return &corev1.PodDNSConfig{
		Nameservers: p.NetWorking.DnsConfig.Nameservers,
	}
}

// DNSPolicy转换
func (i *impl) getK8sDNSPolicy(p *pod.Pod) corev1.DNSPolicy {
	var k8sDnsPolicy corev1.DNSPolicy
	if p.NetWorking == nil {
		return k8sDnsPolicy
	}
	k8sDnsPolicy = corev1.DNSPolicy(p.NetWorking.DnsPolicy)
	return k8sDnsPolicy
}

// HostAliases转换
func (i *impl) getK8sHostAlias(p *pod.Pod) []corev1.HostAlias {
	podK8sHostAliases := make([]corev1.HostAlias, 0)
	if p.NetWorking == nil {
		return podK8sHostAliases
	}
	for _, item := range p.NetWorking.HostAliases {
		podK8sHostAliases = append(podK8sHostAliases, corev1.HostAlias{
			IP:        item.Key,
			Hostnames: strings.Split(item.Value, ","),
		})
	}
	return podK8sHostAliases
}

// Volumes转换
func (i *impl) getK8sVolumes(p *pod.Pod) []corev1.Volume {
	podK8sVolumes := make([]corev1.Volume, 0)
	for _, volume := range p.Volumes {
		source := corev1.VolumeSource{}
		switch volume.Type {
		case VOLUME_EMPTYDIR:
			source = corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			}
		case VOLUME_HOSTPATH:
			pathType := volume.HostPathVolume.Type
			source = corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Type: (*corev1.HostPathType)(&pathType),
					Path: volume.HostPathVolume.Path,
				},
			}
		case VOLUME_CONFIGMAP:
			optional := volume.ConfigMapRefVolume.Optional
			source = corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: volume.ConfigMapRefVolume.Name,
					},
					Optional: &optional,
				},
			}
		case VOLUME_SECRET:
			optional := volume.SecretRefVolume.Optional
			source = corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: volume.SecretRefVolume.Name,
					Optional:   &optional,
				},
			}
		case VOLUME_DOWNWARDAPI:
			items := make([]corev1.DownwardAPIVolumeFile, 0)
			for _, item := range volume.DownwardAPIVolume.Items {
				items = append(items, corev1.DownwardAPIVolumeFile{
					Path: item.Path,
					FieldRef: &corev1.ObjectFieldSelector{
						FieldPath: item.FieldRefPath,
					},
				})
			}
			source = corev1.VolumeSource{
				DownwardAPI: &corev1.DownwardAPIVolumeSource{
					Items: items,
				},
			}
		case VOLUME_PVC:
			source = corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: volume.PVCVolume.Name,
				},
			}
		default:
			continue
		}
		podK8sVolumes = append(podK8sVolumes, corev1.Volume{
			VolumeSource: source,
			Name:         volume.Name,
		})
	}
	return podK8sVolumes
}

// InitContainer结构体转换
func (i *impl) getK8sContainer(podReqContainers []*pod.Container) []corev1.Container {
	podK8sContainers := make([]corev1.Container, 0)
	for _, item := range podReqContainers {
		podK8sContainers = append(podK8sContainers, corev1.Container{
			Name:            item.Name,
			Image:           item.Image,
			ImagePullPolicy: corev1.PullPolicy(item.ImagePullPolicy),
			TTY:             item.Tty,
			Command:         item.Command,
			Args:            item.Args,
			WorkingDir:      item.WorkingDir,
			SecurityContext: &corev1.SecurityContext{
				Privileged: &item.Privileged,
			},
			Ports:          i.getK8sPorts(item.Ports),
			Env:            i.getK8sEnv(item.Envs),
			EnvFrom:        i.getK8sEnvFrom(item.EnvsFrom),
			VolumeMounts:   i.getK8sVolumeMounts(item.VolumeMounts),
			StartupProbe:   i.getK8sContainerProbe(item.StartupProbe),
			LivenessProbe:  i.getK8sContainerProbe(item.LivenessProbe),
			ReadinessProbe: i.getK8sContainerProbe(item.ReadinessProbe),
			Resources:      i.getK8sResources(item.Resources),
		})
	}
	return podK8sContainers
}

// Resources转换
func (i *impl) getK8sResources(podReqResources *pod.Resources) corev1.ResourceRequirements {
	var k8sPodResources corev1.ResourceRequirements
	if podReqResources == nil {
		return k8sPodResources
	}
	if !podReqResources.Enable {
		return k8sPodResources
	}
	k8sPodResources.Requests = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResources.CpuRequest)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResources.MemRequest)) + "Mi"),
	}
	k8sPodResources.Limits = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResources.CpuLimit)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResources.MemLimit)) + "Mi"),
	}
	return k8sPodResources
}

// StartupProbe转换
func (i *impl) getK8sContainerProbe(podReqProbe *pod.ContainerProbe) *corev1.Probe {
	if podReqProbe == nil {
		return nil
	}
	if !podReqProbe.Enable {
		return nil
	}
	var k8sProbe corev1.Probe
	switch podReqProbe.Type {
	case PROBE_HTTP:
		httpGet := podReqProbe.HttpGet
		k8sHttpHeaders := make([]corev1.HTTPHeader, 0)
		for _, header := range httpGet.HttpHeaders {
			k8sHttpHeaders = append(k8sHttpHeaders, corev1.HTTPHeader{
				Name:  header.Key,
				Value: header.Value,
			})
		}
		k8sProbe.HTTPGet = &corev1.HTTPGetAction{
			Scheme:      corev1.URIScheme(httpGet.Scheme),
			Host:        httpGet.Host,
			Port:        intstr.FromInt(int(httpGet.Port)),
			Path:        httpGet.Path,
			HTTPHeaders: k8sHttpHeaders,
		}
		if podReqProbe.ProbeTime != nil {
			k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
			k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
			k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
			k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
			k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
		}
	case PROBE_TCP:
		tcpSocket := podReqProbe.TcpSocket
		k8sProbe.TCPSocket = &corev1.TCPSocketAction{
			Host: tcpSocket.Host,
			Port: intstr.FromInt(int(tcpSocket.Port)),
		}
		if podReqProbe.ProbeTime != nil {
			k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
			k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
			k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
			k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
			k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
		}
	case PROBE_EXEC:
		exec := podReqProbe.Exec
		k8sProbe.Exec = &corev1.ExecAction{
			Command: exec.Command,
		}
		if podReqProbe.ProbeTime != nil {
			k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
			k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
			k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
			k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
			k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
		}
	}
	return &k8sProbe
}

// VolumeMounts转换
func (i *impl) getK8sVolumeMounts(podReqVolumeMounts []*pod.VolumeMounts) []corev1.VolumeMount {
	podK8sVolumeMounts := make([]corev1.VolumeMount, 0)
	for _, mount := range podReqVolumeMounts {
		podK8sVolumeMounts = append(podK8sVolumeMounts, corev1.VolumeMount{
			Name:      mount.MountName,
			MountPath: mount.MountPath,
			ReadOnly:  mount.ReadOnly,
		})
	}
	return podK8sVolumeMounts
}

// EnvFrom转换
func (i *impl) getK8sEnvFrom(podReqEnvsFrom []*pod.EnvVarFromResource) []corev1.EnvFromSource {
	podK8sEnvsFrom := make([]corev1.EnvFromSource, 0)
	for _, fromResource := range podReqEnvsFrom {
		// 前缀通用
		envFromResource := corev1.EnvFromSource{
			Prefix: fromResource.Prefix,
		}
		switch fromResource.RefType {
		case REF_TYPE_CONFIGMAP:
			envFromResource.ConfigMapRef = &corev1.ConfigMapEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			podK8sEnvsFrom = append(podK8sEnvsFrom, envFromResource)
		case REF_TYPE_SECRET:
			envFromResource.SecretRef = &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			podK8sEnvsFrom = append(podK8sEnvsFrom, envFromResource)
		}
	}
	return podK8sEnvsFrom
}

// Env转换
func (i *impl) getK8sEnv(podReqEnvs []*pod.EnvVar) []corev1.EnvVar {
	podK8sEnvs := make([]corev1.EnvVar, 0)
	for _, item := range podReqEnvs {
		envVar := corev1.EnvVar{
			Name: item.Name,
		}
		switch item.Type {
		case REF_TYPE_CONFIGMAP:
			envVar.ValueFrom = &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					Key: item.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: item.RefName,
					},
				},
			}
			podK8sEnvs = append(podK8sEnvs, envVar)
		case REF_TYPE_SECRET:
			envVar.ValueFrom = &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: item.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: item.RefName,
					},
				},
			}
			podK8sEnvs = append(podK8sEnvs, envVar)
		default:
			envVar.Value = item.Value
			podK8sEnvs = append(podK8sEnvs, envVar)
		}
	}
	return podK8sEnvs
}

// Ports转换
func (i *impl) getK8sPorts(podReqPorts []*pod.ContainerPort) []corev1.ContainerPort {
	podK8sContainerPorts := make([]corev1.ContainerPort, 0)
	for _, item := range podReqPorts {
		podK8sContainerPorts = append(podK8sContainerPorts, corev1.ContainerPort{
			Name:          item.Name,
			HostPort:      item.HostPort,
			ContainerPort: item.ContainerPort,
		})
	}
	return podK8sContainerPorts
}

// 污点转换
func (i *impl) getK8sTolerations(p *pod.Pod) []corev1.Toleration {
	podK8sTolerations := make([]corev1.Toleration, 0)
	for _, item := range p.Tolerations {
		podK8sTolerations = append(podK8sTolerations, corev1.Toleration{
			Key:      item.Key,
			Value:    item.Value,
			Operator: corev1.TolerationOperator(item.Operator),
		})
	}
	return podK8sTolerations
}

// 标签转换
func (i *impl) getK8sLabels(p *pod.Pod) map[string]string {
	podK8sLabels := make(map[string]string)
	for _, label := range p.Base.Labels {
		podK8sLabels[label.Key] = label.Value
	}
	return podK8sLabels
}

// 获取Affinity,NodeSelector,NodeName
func (i *impl) getNodeK8sScheduling(p *pod.Pod) (affinity *corev1.Affinity, nodeSelector map[string]string, nodeName string) {
	nodeScheduling := p.NodeScheduling
	if nodeScheduling == nil {
		return
	}
	switch nodeScheduling.Type {
	case SCHEDULING_NODENAME:
		nodeName = nodeScheduling.NodeName
		return
	case SCHEDULING_NODESELECTOR:
		nodeSelectorMap := make(map[string]string)
		for _, item := range nodeScheduling.NodeSelector {
			nodeSelectorMap[item.Key] = item.Value
		}
		nodeSelector = nodeSelectorMap
		return
	case SCHEDULING_NODEAFFINITY:
		nodeSelectorTermExpressions := nodeScheduling.NodeAffinity
		matchExression := make([]corev1.NodeSelectorRequirement, 0)
		for _, expression := range nodeSelectorTermExpressions {
			matchExression = append(matchExression, corev1.NodeSelectorRequirement{
				Key:      expression.Key,
				Values:   strings.Split(expression.Value, ","),
				Operator: corev1.NodeSelectorOperator(expression.Operator),
			})
		}
		affinity = &corev1.Affinity{
			NodeAffinity: &corev1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{
							MatchExpressions: matchExression,
						},
					},
				},
			},
		}
	case SCHEDULING_NODEANY:
		// do nothing
	default:
		// do nothing
	}
	return
}
