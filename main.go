package qtypes

import (
	"time"
	"github.com/qnib/qframe-utils"
	"github.com/qnib/QNIBCollect/src/github.com/hashicorp/go-version"
)



type QMsg struct {
	qmsgVersion string				`json:qmsg_version`
	Type        string        		`json:"type"`
	Source      string        		`json:"source"`
	SourceID    uint64        		`json:"source_id"`
	Host        string        		`json:"host"`
	Msg         string        		`json:"short_message"`
	Time        time.Time    	 	`json:"time"`
	TimeNano    int64        	 	`json:"time_nano"`
	Level       int          	 	`json:"level"` //https://en.wikipedia.org/wiki/Syslog#Severity_level
	KV          map[string]string	`json:"kv"`
	Data 	    interface{}			`json:"data"`
}

func NewQMsg(typ, source string) QMsg {
	now := time.Now()
	return QMsg{
		qmsgVersion: "0.1.0",
		Type: typ,
		Level: 6,
		Source: source,
		SourceID: qutils.GetGID(),
		Time: now,
		TimeNano: now.UnixNano(),
	}
}

func (qm *QMsg) TimeString() (lout string) {
	return qm.Time.Format("2006-01-02T15:04:05.999999")

}

func (qm *QMsg) LogString() (lout string) {
	switch qm.Level {
	case 0:
		lout = "EMERG"
	case 1:
		lout = "ALERT"
	case 2:
		lout = "CRIT"
	case 3:
		lout = "ERROR"
	case 4:
		lout = "WARN"
	case 5:
		lout = "NOTICE"
	case 6:
		lout = "INFO"
	case 7:
		lout = "DEBUG"
	}
	return
}
