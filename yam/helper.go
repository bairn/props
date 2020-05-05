package yam

import (
	"github.com/prometheus/common/log"
	"github.com/bairn/props/kvs"
	"strings"
)

func ByYaml(content string) *kvs.MapProperties {
	y := NewYamlProperties()
	err := y.Load(strings.NewReader(content))
	if err != nil {
		log.Error(err)
		return nil
	}
	return &y.MapProperties
}
