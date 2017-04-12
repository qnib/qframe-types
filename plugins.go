package qtypes

import (
	"github.com/zpatrick/go-config"
)

type Plugin struct {
	QChan QChan
	Cfg config.Config
	Version string
	Name string
}

func NewPlugin(qChan QChan, cfg config.Config) Plugin {
	return Plugin{
		QChan: qChan,
		Cfg: cfg,
	}
}
