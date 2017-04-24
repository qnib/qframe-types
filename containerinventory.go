package qtypes


import (
	"context"
	"fmt"
	"errors"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
)

const (
	mobyAPI = "v1.29"
)

type ContainerInventory struct {
	Client *client.Client
	Data   map[string]types.ContainerJSON
	IDtoIP map[string]string
}

func NewContainerInventory(cli *client.Client) ContainerInventory {
	return ContainerInventory{
		Client: cli,
		Data: map[string]types.ContainerJSON{},
		IDtoIP: map[string]string{},
	}
}

func NewPlainContainerInventory() ContainerInventory {
	return ContainerInventory{
		Data: map[string]types.ContainerJSON{},
	}
}

func NewContainerInventoryHost(dockerHost string) ContainerInventory {
	engineCli, _ := client.NewClient(dockerHost, mobyAPI, nil, nil)
	return ContainerInventory{
		Client: engineCli,
		Data: map[string]types.ContainerJSON{},
		IDtoIP: map[string]string{},
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

func (ci *ContainerInventory) GetCntByEvent(event events.Message) (cnt types.ContainerJSON, err error) {
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
		cnt, err = ci.Client.ContainerInspect(context.Background(), id)
		ci.Data[id] = cnt
		if cnt.State.Running {
			for _, v := range cnt.NetworkSettings.Networks {
				ci.IDtoIP[id] = v.IPAddress
			}
		} else {
			fmt.Printf("cnt.State: %v\n", cnt.State)
		}
	default:
		cnt, err = ci.Client.ContainerInspect(context.Background(), id)
	}
	return cnt, err
}