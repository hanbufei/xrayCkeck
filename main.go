package main

import (
	"flag"
	"fmt"
	"github.com/hanbufei/xrayCkeck/poc"
)

var (
	proxy  string
	target string
	s      string
	p      string
)

func init() {
	flag.StringVar(&proxy, "proxy", "", "http代理，如http://127.0.0.1:8080")
	flag.StringVar(&target, "target", "", "目标")
	flag.StringVar(&p, "poc", "", "指定yml文件的目录，如果不指定，则采用内置目录")
	flag.StringVar(&s, "s", "", "poc前缀名")
	flag.Parse()
	fmt.Println(`usage: xrayCheck -target http://127.0.0.1 -poc tomcat -path /Users/xrayFiles`)
}
func main() {
	if proxy != "" {
		poc.HttpProxy = proxy
	}
	poc.XrayCheck(target, s, p)
}
