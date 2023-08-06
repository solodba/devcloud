package conf

import (
	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// 初始化Prometheus客户端
func (p *Prometheus) GetPromClient() (promapi.Client, error) {
	client, err := promapi.NewClient(promapi.Config{
		Address: p.Addr(),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 获取Promtheus API
func (p *Prometheus) GetPromAPI() (promv1.API, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.promAPI == nil {
		client, err := p.GetPromClient()
		if err != nil {
			return nil, err
		}
		promAPI := promv1.NewAPI(client)
		p.promAPI = promAPI
	}
	return p.promAPI, nil
}
