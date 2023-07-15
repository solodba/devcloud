package pod

import (
	"strconv"
	"strings"

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
	PROB_EXEC  = "exec"
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
func (p *Pod) PodReq2K8s() *corev1.Pod {
	nodeAffinity, nodeSelector, nodeName := p.getNodeK8sScheduling()
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Base.Name,
			Namespace: p.Base.Namespace,
			Labels:    p.getK8sLabels(),
		},
		Spec: corev1.PodSpec{
			NodeName:       nodeName,
			NodeSelector:   nodeSelector,
			Affinity:       nodeAffinity,
			Tolerations:    p.getK8sTolerations(),
			InitContainers: p.getK8sContainer(p.InitContainers),
			Containers:     p.getK8sContainer(p.Containers),
			Volumes:        p.getK8sVolumes(),
			DNSConfig: &corev1.PodDNSConfig{
				Nameservers: p.NetWorking.DnsConfig.Nameservers,
			},
			DNSPolicy:     corev1.DNSPolicy(p.NetWorking.DnsPolicy),
			HostAliases:   p.getK8sHostAlias(),
			Hostname:      p.NetWorking.HostName,
			RestartPolicy: corev1.RestartPolicy(p.Base.RestartPolicy),
		},
	}
}

// HostAliases转换
func (p *Pod) getK8sHostAlias() []corev1.HostAlias {
	podK8sHostAliases := make([]corev1.HostAlias, 0)
	for _, item := range p.NetWorking.HostAliases {
		podK8sHostAliases = append(podK8sHostAliases, corev1.HostAlias{
			IP:        item.Key,
			Hostnames: strings.Split(item.Value, ","),
		})
	}
	return podK8sHostAliases
}

// Volumes转换
func (p *Pod) getK8sVolumes() []corev1.Volume {
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
func (p *Pod) getK8sContainer(podReqContainers []*Container) []corev1.Container {
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
			Ports:          p.getK8sPorts(item.Ports),
			Env:            p.getK8sEnv(item.Envs),
			EnvFrom:        p.getK8sEnvFrom(item.EnvsFrom),
			VolumeMounts:   p.getK8sVolumeMounts(item.VolumeMounts),
			StartupProbe:   p.getK8sContainerProbe(item.StartupProbe),
			LivenessProbe:  p.getK8sContainerProbe(item.LivenessProbe),
			ReadinessProbe: p.getK8sContainerProbe(item.ReadinessProbe),
			Resources:      p.getK8sResources(item.Resources),
		})
	}
	return podK8sContainers
}

// Resources转换
func (p *Pod) getK8sResources(podReqResources *Resources) corev1.ResourceRequirements {
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
func (p *Pod) getK8sContainerProbe(podReqProbe *ContainerProbe) *corev1.Probe {
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
		k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
		k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
		k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
		k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
		k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
	case PROBE_TCP:
		tcpSocket := podReqProbe.TcpSocket
		k8sProbe.TCPSocket = &corev1.TCPSocketAction{
			Host: tcpSocket.Host,
			Port: intstr.FromInt(int(tcpSocket.Port)),
		}
		k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
		k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
		k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
		k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
		k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
	case PROB_EXEC:
		exec := podReqProbe.Exec
		k8sProbe.Exec = &corev1.ExecAction{
			Command: exec.Command,
		}
		k8sProbe.InitialDelaySeconds = podReqProbe.ProbeTime.InitialDelaySeconds
		k8sProbe.PeriodSeconds = podReqProbe.ProbeTime.PeriodSeconds
		k8sProbe.TimeoutSeconds = podReqProbe.ProbeTime.TimeoutSeconds
		k8sProbe.SuccessThreshold = podReqProbe.ProbeTime.SuccessThreshold
		k8sProbe.FailureThreshold = podReqProbe.ProbeTime.FailureThreshold
	}
	return &k8sProbe
}

// VolumeMounts转换
func (p *Pod) getK8sVolumeMounts(podReqVolumeMounts []*VolumeMounts) []corev1.VolumeMount {
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
func (p *Pod) getK8sEnvFrom(podReqEnvsFrom []*EnvVarFromResource) []corev1.EnvFromSource {
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
func (p *Pod) getK8sEnv(podReqEnvs []*EnvVar) []corev1.EnvVar {
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
func (p *Pod) getK8sPorts(podReqPorts []*ContainerPort) []corev1.ContainerPort {
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
func (p *Pod) getK8sTolerations() []corev1.Toleration {
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
func (p *Pod) getK8sLabels() map[string]string {
	podK8sLabels := make(map[string]string)
	for _, label := range p.Base.Labels {
		podK8sLabels[label.Key] = label.Value
	}
	return podK8sLabels
}

// 获取Affinity,NodeSelector,NodeName
func (p *Pod) getNodeK8sScheduling() (affinity *corev1.Affinity, nodeSelector map[string]string, nodeName string) {
	nodeScheduling := p.NodeScheduling
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
