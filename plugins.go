package qtypes

import (
	"log"
	"fmt"
	"strings"
	"github.com/zpatrick/go-config"
)

type Plugin struct {
	QChan 	QChan
	Cfg 	config.Config
	Typ		string
	Version string
	Name 	string
}

func NewPlugin(qChan QChan, cfg config.Config) Plugin {
	return Plugin{
		QChan: qChan,
		Cfg: cfg,
	}
}

func logStrToInt(level string) int {
	def := 6
	switch level {
	case "error":
		return 3
	case "warn":
		return 4
	case "notice":
		return 5
	case "info":
		return 6
	case "debug":
		return 7
	default:
		return def
	}
}

func (p *Plugin) CfgStringOr(path, alt string) string {
	res, _ := p.Cfg.StringOr(fmt.Sprintf("%s.%s.host", p.Typ, p.Name), alt)
	return res
}

func (p *Plugin) Log(logLevel, msg string) {
	dL, _ := p.Cfg.StringOr("log.level", "info")
	dI := logStrToInt(dL)
	lI := logStrToInt(logLevel)
	if dI >= lI {
		log.Printf("[%+6s] %s >> %s", strings.ToUpper(logLevel), p.Name, msg)
	}
}

func NewNamedPlugin(qChan QChan, cfg config.Config, typ, name, version string) Plugin {
	p := Plugin{
		QChan: qChan,
		Cfg:   cfg,
	}
	p.Typ = typ
	p.Version = version
	p.Name = name
	return p
}
