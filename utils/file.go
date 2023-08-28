package utils

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/hanbufei/xrayCkeck/pkg/xray/structs"
	"gopkg.in/yaml.v2"
)

func LoadMultiPocFromPath(path string, pocname string) []*structs.Poc {
	var pocs []*structs.Poc
	poclist := path + gfile.Separator + pocname + "-*.yml"
	matchNames, _ := gfile.Glob(poclist, true)
	for _, f := range matchNames {
		p := &structs.Poc{}
		yamlFile := gfile.GetContents(path + gfile.Separator + f)
		yaml.Unmarshal([]byte(yamlFile), p)
		pocs = append(pocs, p)
	}
	return pocs
}
