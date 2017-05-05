package qtypes


import (
	"fmt"
	"strings"
	"github.com/docker/docker/api/types"
)


// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/net/helper.go
type NetStats struct {
	Base
	Container     *types.Container
	NameInterface string
	RxBytes       float64
	RxDropped     float64
	RxErrors      float64
	RxPackets     float64
	TxBytes       float64
	TxDropped     float64
	TxErrors      float64
	TxPackets     float64
}


func (ns *NetStats) ToMetrics(src string) []Metric {
	dim := map[string]string{
		"container_id": ns.Container.ID,
		"container_name": strings.Trim(ns.Container.Names[0], "/"),
		"image_name": ns.Container.Image,
		"command": strings.Replace(ns.Container.Command, " ", "#", -1),
		"created": fmt.Sprintf("%d", ns.Container.Created),
	}
	for k, v := range ns.Container.Labels {
		dim[k] = strings.Replace(v, " ", "#", -1)
	}
	return []Metric{
		ns.NewExtMetric(src, "network.total.rx.bytes", Gauge, ns.RxBytes, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.rx.dropped", Gauge, ns.RxDropped, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.rx.errors", Gauge, ns.RxErrors, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.rx.packets", Gauge, ns.RxPackets, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.tx.bytes", Gauge, ns.TxBytes, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.tx.dropped", Gauge, ns.TxDropped, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.tx.errors", Gauge, ns.TxErrors, dim, ns.Time, true),
		ns.NewExtMetric(src, "network.total.tx.packets", Gauge, ns.TxPackets, dim, ns.Time, true),
	}
}
