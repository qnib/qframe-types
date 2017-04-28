package qtypes


import (
	"fmt"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
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
	err = errors.New(fmt.Sprintf("No container found with key '%s'", key))
	return
}

func (ci *ContainerInventory) GetCntByID(id string) (cnt types.ContainerJSON, err error) {
	if cnt, ok := ci.Data[id];ok {
		return cnt, err
	}
	return cnt, err
}

func (ci *ContainerInventory) GetCntByIP(ip string) (cnt types.ContainerJSON, err error) {
	for id, v := range ci.IDtoIP {
		if ip == v {
			return ci.GetCntByID(id)
		}

	}
	return cnt, err
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
	}
	return err
}

func (ci *ContainerInventory) GetCntByEvent(event events.Message) (cnt types.ContainerJSON, err error) {
	fmt.Println("Sorry, have to redo GetCntByEvent...")
	return
	id := event.Actor.ID
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
		ci.Data[id] = cnt
		if cnt.State.Running {
			for _, v := range cnt.NetworkSettings.Networks {
				ci.IDtoIP[id] = v.IPAddress
			}
		} else {
			fmt.Printf("cnt.State: %v\n", cnt.State)
		}
	}
	return cnt, err
}