package cert

import (
	"github.com/limeschool/easy-admin/server/global"
	"io"
	"os"
)

func Init() {
	confList := global.Config.Cert
	cert := make(map[string][]byte)
	for _, item := range confList {
		file, err := os.Open(item.Path)
		if err != nil {
			panic("cert初始化失败:" + err.Error())
		}
		key, err := io.ReadAll(file)
		if err != nil {
			panic("读取cert证书失败:" + err.Error())
		}
		cert[item.Name] = key
	}
	global.Cert = cert
}
