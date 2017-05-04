package qtypes

/*
import (
	"time"

	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
)


// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/net/helper.go

///////////////// NetRaw

type NetRaw struct {
	Time      time.Time
	RxBytes   uint64
	RxDropped uint64
	RxErrors  uint64
	RxPackets uint64
	TxBytes   uint64
	TxDropped uint64
	TxErrors  uint64
	TxPackets uint64
}

func createNetRaw(time time.Time, stats *dc.NetworkStats) NetRaw {
	return NetRaw{
		Time:      time,
		RxBytes:   stats.RxBytes,
		RxDropped: stats.RxDropped,
		RxErrors:  stats.RxErrors,
		RxPackets: stats.RxPackets,
		TxBytes:   stats.TxBytes,
		TxDropped: stats.TxDropped,
		TxErrors:  stats.TxErrors,
		TxPackets: stats.TxPackets,
	}
}

///////////////// NetStats


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


type NetStats struct {
	Time          time.Time
	Container     *docker.Container

}

func GetNetworkStatsPerContainer(tats dc.Stat) []NetStats {
	formattedStats := []NetStats{}
	for _, myStats := range rawStats {
		for nameInterface, rawnNetStats := range myStats.Stats.Networks {
			formattedStats = append(formattedStats, GetNetworkStats(nameInterface, &rawnNetStats, &myStats))
		}
	}

	return formattedStats
}



func createNetRaw(time time.Time, stats *dc.NetworkStats) NetRaw {
	return NetRaw{
		Time:      time,
		RxBytes:   stats.RxBytes,
		RxDropped: stats.RxDropped,
		RxErrors:  stats.RxErrors,
		RxPackets: stats.RxPackets,
		TxBytes:   stats.TxBytes,
		TxDropped: stats.TxDropped,
		TxErrors:  stats.TxErrors,
		TxPackets: stats.TxPackets,
	}
}

func (n *NetService) checkStats(containerID string, nameInterface string) bool {
	if _, exist := n.NetworkStatPerContainer[containerID][nameInterface]; exist {
		return true
	}
	return false
}

func (n *NetService) getRxBytesPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.RxBytes, newStats.RxBytes)
}

func (n *NetService) getRxDroppedPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.RxDropped, newStats.RxDropped)
}

func (n *NetService) getRxErrorsPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.RxErrors, newStats.RxErrors)
}

func (n *NetService) getRxPacketsPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.RxPackets, newStats.RxPackets)
}

func (n *NetService) getTxBytesPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.TxBytes, newStats.TxBytes)
}

func (n *NetService) getTxDroppedPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.TxDropped, newStats.TxDropped)
}

func (n *NetService) getTxErrorsPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.TxErrors, newStats.TxErrors)
}

func (n *NetService) getTxPacketsPerSecond(newStats *NetRaw, oldStats *NetRaw) float64 {
	duration := newStats.Time.Sub(oldStats.Time)
	return n.calculatePerSecond(duration, oldStats.TxPackets, newStats.TxPackets)
}

func (n *NetService) calculatePerSecond(duration time.Duration, oldValue uint64, newValue uint64) float64 {
	value := float64(newValue) - float64(oldValue)
	if value < 0 {
		value = 0
	}
	return value / duration.Seconds()
}

*/
