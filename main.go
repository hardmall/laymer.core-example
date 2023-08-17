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
	database.SetService(dsvc) //设置查询

	q := database.DService{DbName: "xdb"} //使用xdb查询
	var dd []database.APIQueryApiInfo
	database.Svc().DB().Where(&database.APIQueryApiInfo{}).Find(&dd)
	a, b, c, err := q.QueryServiceWithPage(dd, make(map[string]interface{}), 2, 5)
	if err != nil {
		log.Fatal("查询发生错误", err)
	}
	log.Println(a, b, c)

	//使用缓存查询
	info := database.DeCheck().GetTableByID(1)
	log.Println(info)
}
