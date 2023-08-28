package structs

import (
	"github.com/hanbufei/xrayCkeck/pkg/xray/structs"
	"net/http"
)

var (
	ReversePlatformType      structs.ReverseType
	DnslogCNGetDomainRequest *http.Request
	DnslogCNGetRecordRequest *http.Request
)

func InitReversePlatform() {
	ReversePlatformType = structs.ReverseType_DnslogCN
	// 设置请求相关参数
	DnslogCNGetDomainRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getdomain.php", nil)
	DnslogCNGetRecordRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getrecords.php", nil)
}
