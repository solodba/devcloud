package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// prometheus收集
func (h *handler) GetMetrics(r *restful.Request, w *restful.Response) {
	handler := promhttp.Handler()
	handler.ServeHTTP(w.ResponseWriter, r.Request)
}
