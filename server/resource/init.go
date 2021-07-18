package resource

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"
)

//go:embed wxh5/*
var Wxh5 embed.FS

//go:embed console/*
var Console embed.FS

//go:embed www/*
var WebsiteTemplete embed.FS

//go:embed favicon.ico
var FaviconIco embed.FS

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
		file, err := assets.Open(assetPath)
		log.Default().Println(assetPath)
		if err != nil {

			return nil, err
		}

		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

//模拟nginx 反向代理 tryfile
func TryfileHandler(prefix string, assets embed.FS, root string, dstfile string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)
		file, err := assets.Open(assetPath)
		//如果资源存在,那么直接处理文件
		if err == nil {
			//log.Default().Println(err.Error())
			return file, err
		}
		assetPath = path.Join(root, dstfile)
		file, err = assets.Open(assetPath)
		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})
	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

func TempleteFs(assets embed.FS, root string) fs.FS {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)
		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}
		return file, err
	})
	return handler
}
