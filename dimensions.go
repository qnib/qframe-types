package qtypes

import (
	"fmt"
	"strings"
)

type Dimensions struct {
	Map map[string]string
}

func NewDimensions() Dimensions {
	return Dimensions{
		Map: map[string]string{},
	}
}

func (dim *Dimensions) Add(key,val string) {
	dim.Map[key] = val
}

func (dim *Dimensions) String() string {
	res := []string{}
	for k,v := range dim.Map {
		res = append(res, fmt.Sprintf("%s=%s", k,v))
	}
	return strings.Join(res, ",")
}