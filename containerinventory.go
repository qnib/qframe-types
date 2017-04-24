package qtypes


import (
	"context"
	"fmt"
	"log"
	"errors"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

type ContainerInventory struct {
	Client *client.Client
	Data   map[string]types.Container
}

func NewContainerInventory(cli *client.Client) ContainerInventory {
	return ContainerInventory{
		Client: cli,
		Data: map[string]types.Container{},
	}
}
func (ci *ContainerInventory) SetCnt(key string, cnt types.Container) (err error) {
	ci.Data[key] = cnt
	return
}

func (ci *ContainerInventory) GetCnt(key string) (cnt types.Container, err error) {
	if cnt, ok := ci.Data[key];ok {
		return cnt, err
	}
	err = errors.New(fmt.Sprintf("No container found with key '%s'", key))
	return
}


func (ci *ContainerInventory) GetCntByID(id string) (cnt types.Container, err error) {
	if cnt, ok := ci.Data[id];ok {
		return cnt, err
	} else {
		args := filters.NewArgs()
		args, err = filters.ParseFlag(fmt.Sprintf("id=%s",id), args)
		df := types.ContainerListOptions{Filters: args, All: true}
		cnts, _ := ci.Client.ContainerList(context.Background(), df)
		if len(cnts) == 1 {
			cnt = cnts[0]
			return cnt, err
		}
	}
	return cnt, err
}