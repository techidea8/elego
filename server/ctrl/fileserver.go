package ctrl

import (
	"net/http"
	"turingapp/resource"
)

//注册文件或文件夹服务器
func RegisterFileServer() {
	http.Handle("/console/", resource.TryfileHandler("/console/", resource.Console, "console/", "index.html"))
	http.Handle("/wxh5/", resource.TryfileHandler("/wxh5/", resource.Wxh5, "wxh5/", "index.html"))
	http.Handle("/favicon.ico", resource.AssetHandler("/", resource.FaviconIco, "/"))
	//http.Handle("/", resource.TempleteFs(resource.WebsiteTemplete, "www/"))
}
