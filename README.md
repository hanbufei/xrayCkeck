# xrayCkeck
 导入xray的yml文件，进行检查的工具
# 当作api导入
```bigquery
import "github.com/hanbufei/xrayCkeck/poc"
func main() {
	poc.XrayCheck("http://127.0.0.1","tomcat","")
}
```