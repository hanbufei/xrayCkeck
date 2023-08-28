package poc

import (
	"embed"
	"github.com/hanbufei/xrayCkeck/pkg/check"
	common_structs "github.com/hanbufei/xrayCkeck/pkg/common/structs"
	xray_requests "github.com/hanbufei/xrayCkeck/pkg/xray/requests"
	"github.com/hanbufei/xrayCkeck/pkg/xray/structs"
	"github.com/hanbufei/xrayCkeck/utils"
	"time"
)

//go:embed xrayFiles
var XrayPocs embed.FS

var HttpProxy string

func XrayCheck(target string, pocname string, path string) []string {
	var xrayPocs []*structs.Poc
	if path != "" {
		xrayPocs = utils.LoadMultiPocFromPath(path, pocname)
	} else {
		xrayPocs = utils.LoadMultiPoc(XrayPocs, pocname)
	}
	common_structs.InitReversePlatform()
	_ = xray_requests.InitHttpClient(10, HttpProxy, time.Duration(5)*time.Second)
	xrayTotalReqeusts := 0
	for _, poc := range xrayPocs {
		ruleLens := len(poc.Rules)
		if poc.Transport == "tcp" || poc.Transport == "udp" {
			ruleLens += 1
		}
		xrayTotalReqeusts += 1 * ruleLens
	}
	if xrayTotalReqeusts == 0 {
		xrayTotalReqeusts = 1
	}
	xray_requests.InitCache(xrayTotalReqeusts)
	return check.XrayStart(target, xrayPocs)
}
