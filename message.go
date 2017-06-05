package qtypes

import (
	"strings"
	"github.com/docker/docker/api/types"
	"github.com/qnib/qframe-utils"
)

const (
	MsgCEE = "cee"
	MsgTCP = "tcp"
	MsgDLOG = "docker-log"
	MsgMetric = "metric" //needs to have name,time and value field ; optional tags (key1=val1,key2=val2)
)


type Message struct {
	Base
	Container   types.ContainerJSON
	Name       	string            	`json:"name"`
	LogLevel    string				`json:"loglevel"`
	MessageType	string            	`json:"type"`
	Message     string            	`json:"value"`
	KV			map[string]string 	`json:"data"`
}

func NewMessage(base Base, name, mType, msg string) Message {
	m := Message{
		Base: base,
		Name: name,
		Container: types.ContainerJSON{},
		LogLevel: "INFO",
		MessageType: mType,
		Message: msg,
		KV: map[string]string{},
	}
	m.SourceID = int(qutils.GetGID())
	return m
}

func NewContainerMessage(base Base, cnt types.ContainerJSON, name, mType, msg string) Message {
	m := NewMessage(base, name, mType, msg)
	m.Container = cnt
	return m
}

func (m *Message) GetContainerName() string {
	if m.Container.Name != "" {
		return strings.Trim(m.Container.Name, "/")
	} else {
		return "<none>"
	}
}
