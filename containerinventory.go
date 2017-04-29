package qtypes


import (
	"fmt"
	"errors"
	"github.com/docker/docker/api/types"
	"strings"
)

type ContainerInventory struct {
	Data   map[string]types.ContainerJSON
	IDtoIP map[string]string
}

func NewContainerInventory() ContainerInventory {
	return ContainerInventory{
		Data: map[string]types.ContainerJSON{},
		IDtoIP: map[string]string{},
	}
}

func NewPlainContainerInventory() ContainerInventory {
	return ContainerInventory{
		Data: map[string]types.ContainerJSON{},
	}
}

func (ci *ContainerInventory) SetCnt(key string, cnt types.ContainerJSON) (err error) {
	ci.Data[key] = cnt
	return
}

func (ci *ContainerInventory) GetCnt(key string) (cnt types.ContainerJSON, err error) {
	if cnt, ok := ci.Data[key];ok {
		return cnt, err
	}
	return cnt, errors.New(fmt.Sprintf("No container found with key '%s'", key))
}

func (ci *ContainerInventory) GetCntByID(id string) (cnt types.ContainerJSON, err error) {
	return ci.GetCnt(id)
}

func (ci *ContainerInventory) GetCntByIP(ip string) (cnt types.ContainerJSON, err error) {
	fmt.Printf("Check for IP %s\n", ip)
	for id, v := range ci.IDtoIP {
		if ip == v {
			return ci.GetCntByID(id)
		}

	}
	return cnt, errors.New(fmt.Sprintf("No container found with IP '%s'", ip))
}

func (ci *ContainerInventory) SetCntByEvent(ce ContainerEvent) (err error) {
	id := ce.Event.Actor.ID
	event := ce.Event
	if event.Type != "container" {
		return
	}
	switch event.Action {
	case "die", "destroy":
		if _, ok := ci.IDtoIP[id]; ok {
			delete(ci.IDtoIP, id)
		}
		if _, ok := ci.Data[id]; ok {
			delete(ci.Data, id)
		}
		return
	case "start":
		cnt := ce.Container
		if err != nil {
			return err
		}
		ci.Data[id] = cnt
		if cnt.State.Running {
			for _, v := range cnt.NetworkSettings.Networks {
				ci.IDtoIP[id] = v.IPAddress
			}
		} else {
			fmt.Printf("cnt.State: %v\n", cnt.State)
		}
	default:
		if strings.HasPrefix(event.Action, "exec_") {
			// Not sure what to do here... it happens so often...
			return
		}
	}
	return err
}

func (ci *ContainerInventory) GetCntByEvent(ce ContainerEvent) (cnt types.ContainerJSON, err error) {
	id := ce.Event.Actor.ID
	cnt = ce.Container
	if ce.Event.Type != "container" {
		return
	}
	switch ce.Event.Action {
	case "die", "destroy":
		if _, ok := ci.IDtoIP[id]; ok {
			delete(ci.IDtoIP, id)
		}
		if _, ok := ci.Data[id]; ok {
			delete(ci.Data, id)
		}
		return
	case "start":
		ci.Data[id] = cnt
		if cnt.State.Running {
			for _, v := range cnt.NetworkSettings.Networks {
				ci.IDtoIP[id] = v.IPAddress
			}
		}
	}
	return cnt, err
}