package main

import (
	"net/url"
	"os"

	"github.com/laymer110/laymer.core/database"
)

var (
	local = os.DirFS("./home")
)

func main() {
	dsvc := database.CreateServiceBaseFS(local, &database.BaseFSConfig{
		AppName:       "simple",
		WebName:       "dbapp",
		Version:       "1.0.0",
		WebConfigPath: "wconfig.json",
		GDBConfigPath: "gdb.json",
		XDBConfigPath: "xdb.json",
	})
	database.SetService(dsvc)
	q := database.DService{DbName: ""}
	var dd []database.APIQueryApiInfo
	database.Svc().DB().Where(&database.APIQueryApiInfo{}).Find(&dd)
	q.QueryService(dd, make(map[string]interface{}), url.Values{"pageNum": []string{"2"}})
}
