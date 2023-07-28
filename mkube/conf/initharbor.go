package conf

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取Harbor CaCerts
func (h *Harbor) GetHarborCaCerts() (string, error) {
	dj, err := ioutil.ReadFile(h.CaCerts)
	if err != nil {
		return "", err
	}
	return string(dj), nil
}

// 初始化Harbor Client
func (h *Harbor) InitHarborClient() (*http.Client, error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.client == nil {
		url := fmt.Sprintf("%s://%s%s", h.Scheme, h.Host, GET_CONFIGURATIONS)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(h.Username, h.Password)
		caCerts, err := h.GetHarborCaCerts()
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(caCerts))
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					RootCAs:            caCertPool,
				},
			},
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resMap := make(map[string]any)
		err = json.Unmarshal(respBytes, &resMap)
		if err != nil {
			return nil, err
		}
		if _, ok := resMap["errors"]; ok {
			return nil, fmt.Errorf("认证失败: %s", string(respBytes))
		}
		h.client = client
	}
	return h.client, nil
}
