package main

import (
	"os"

	"log"

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
		XDBConfigPath: []string{"xdb.json"},
	})
	database.SetService(dsvc)
	q := database.DService{DbName: "xdb"}
	var dd []database.APIQueryApiInfo
	database.Svc().DB().Where(&database.APIQueryApiInfo{}).Find(&dd)
	a, b, c, err := q.QueryServiceWithPage(dd, make(map[string]interface{}), 2, 5)
	if err != nil {
		log.Fatal("查询发生错误", err)
	}
	log.Println(a, b, c)
}
