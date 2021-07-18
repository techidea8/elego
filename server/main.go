package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"time"
	"turingapp/global"
	"turingapp/initialize"

	_ "github.com/dimiro1/banner/autoload"
	"go.uber.org/zap"
)

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// AssetHandler returns an http.Handler that will serve files from
// the Assets embed.FS. When locating a file, it will strip the given
// prefix from the request and prepend the root to the filesystem.
func AssetHandler(prefix string, assets embed.FS, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}

		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}



func main() {

	//初始化环境变量
	initialize.InitConfig()
	//fmt.Printf("%+V", global.Config)
	//初始化日志
	initialize.InitLogger()
	//初始化redis
	initialize.InitRedis()
	//初始化数据库
	initialize.InitDatabase()

	initialize.InitAdmin("15367151352", "rootme@1")

	//初始化路由器
	initialize.InitRouter()




	global.Log.Info("run @", zap.Int("port", global.Conf.System.Port), zap.String("board=", global.Conf.System.Dashboard))
	err := http.ListenAndServe(fmt.Sprintf(":%d", global.Conf.System.Port), nil)
	if err != nil {
		global.Log.Error("", zap.Error(err))
	}
	time.Sleep(5 * time.Second)
}
