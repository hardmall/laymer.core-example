# laymer.core-example

this project will show repository <https://github.com/laymer110/laymer.core.git> how used.

## use database.CreateServiceBaseFS Init

```
	dsvc := database.CreateServiceBaseFS(local, &database.BaseFSConfig{
		AppName:       "simple",
		WebName:       "dbapp",
		Version:       "1.0.0",
		WebConfigPath: "wconfig.json",
		GDBConfigPath: "gdb.json",
		XDBConfigPath: []string{"xdb.json"},
	})
	database.SetService(dsvc) //设置查询
```

## use database.SetService Init

