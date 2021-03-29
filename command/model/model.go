package model

import (
	"fmt"
	"github.com/coder2z/g-saber/xconsole"
	"os"
	"path/filepath"
	"yangon/constant"
	"yangon/tools"

	"github.com/BurntSushi/toml"
	"github.com/coder2z/g-saber/xcfg"
	"github.com/coder2z/g-saber/xflag"
	"github.com/coder2z/g-server/datasource/manager"
	"github.com/coder2z/g-server/xinvoker"
	database "github.com/coder2z/g-server/xinvoker/gorm"
)

type List struct {
	Field   string `gorm:"Field"`
	Type    string `gorm:"Type"`
	Null    string `gorm:"Null"`
	Key     string `gorm:"Key"`
	Default string `gorm:"Default"`
	Extra   string `gorm:"Extra"`
}

//todo 生成handle map
//todo server map
//todo map

func (options *RunOptions) Run() {
	dir, _ := os.Getwd()
	var err error

	data, err := manager.NewDataSource(xflag.NString("go", "xcfg"))
	tools.MustCheck(err)

	err = xcfg.LoadFromDataSource(data, toml.Unmarshal)

	tools.MustCheck(err)

	xinvoker.Register(
		database.Register(options.dbKey),
	)
	err = xinvoker.Init()
	tools.MustCheck(err)
	//链接数据库
	db := database.Invoker(options.dbLabel)

	//拉取模板
	tools.MustCheck(tools.GitClone(constant.GitUrl, filepath.Join(dir, "tmp", options.ProjectName)))
	//defer删除拉取的模板
	defer tools.RemoveAllList("tmp")
	//查找表
	rows, err := db.Raw("show tables;").Rows()
	tools.MustCheck(err)
	defer rows.Close()
	var table string
	for rows.Next() {
		tools.MustCheck(rows.Scan(&table))
		//把表名进行驼峰式转换
		modelName := tools.StrFirstToUpper(table)
		//查字段名
		listRows, err := db.Raw(fmt.Sprintf("show columns from %s;", table)).Rows()
		tools.MustCheck(err)
		var TableFieldList, TableFieldMap string
		imports := ""
		Id := "ID"
		for listRows.Next() {
			list := new(List)
			//查询所有字段
			_ = listRows.Scan(&list.Key, &list.Type, &list.Default, &list.Extra, &list.Field, &list.Null)
			//把字段名进行驼峰式转换
			upper := tools.StrFirstToUpper(list.Key)
			//判断是不是主键
			if tools.IsPRI(list.Type) {
				Id = upper
			}
			var structType string
			var tmpIsTime bool
			//把对应的字段类型转换为结构体类型
			structType, tmpIsTime = tools.SqlType2StructType(list.Type, list.Null)
			//组合结构体中的字段，字符串
			TableFieldList += fmt.Sprintf("%s\t%s\n\t", upper, structType)
			if !tmpIsTime {
				TableFieldMap += fmt.Sprintf("%s\t%s\t `form:\"%s\" json:\"%s\" validate:\"required\"` \n\t", upper, structType, list.Key, list.Key)
			}

			if v, ok := tools.EImportsHead[structType]; ok {
				imports += fmt.Sprintf("%s\n", v)
			}
		}
		// model
		{
			var modelText string
			//获取到模板文件
			//模板替换
			modelText, err = tools.ParseTmplFile(filepath.Join(dir, "tmp", options.ProjectName, "model", "model.go.tmpl"), map[string]string{
				"TableFieldList": TableFieldList,
				"ProjectName":    options.ProjectName,
				"appName":        options.AppName,
				"TableName":      modelName,
				"tableName":      table,
				"ID":             Id,
				"Imports":        imports,
			})
			tools.MustCheck(err)
			//模板替换文件夹位置
			modelPath := filepath.Join(dir, "internal", "{{appName}}", "model", "{{table}}")
			modelPath = tools.ReplaceAllData(modelPath, map[string]string{
				"{{appName}}": options.AppName,
				"{{table}}":   table,
			})
			//创建文件夹
			tools.MustCheck(os.MkdirAll(modelPath, 777))
			//模板替换文件位置
			modelFile := filepath.Join("{{path}}", "{{table}}.go")
			modelFile = tools.ReplaceAllData(modelFile, map[string]string{
				"{{path}}":  modelPath,
				"{{table}}": table,
			})
			//判断文件存在，如果存在 就备份之前文件
			if tools.CheckFileIsExist(modelFile) {
				tools.MustCheck(os.Rename(modelFile, fmt.Sprintf("%s.bak", modelFile)))
			}
			//向文件中写入数据
			tools.WriteToFile(modelFile, modelText)
			xconsole.Greenf("model", modelFile)
		}

		//handle
		{
			var handleText string
			//获取到模板文件
			handleText, err = tools.ParseTmplFile(filepath.Join(dir, "tmp", options.ProjectName, "model", "handle.go.tmpl"), map[string]string{
				"ProjectName": options.ProjectName,
				"appName":     options.AppName,
				"AppName":     tools.StrFirstToUpper(options.AppName),
				"TableName":   modelName,
				"tableName":   table,
			})
			tools.MustCheck(err)
			//模板替换文件位置
			handlePath := filepath.Join(dir, "internal", "{{appName}}", "api", "{{version}}", "handle")
			handlePath = tools.ReplaceAllData(handlePath, map[string]string{
				"{{appName}}": options.AppName,
				"{{version}}": tools.UnStrFirstToUpper(options.Version),
			})
			//创建文件夹
			tools.MustCheck(os.MkdirAll(handlePath, 777))
			handleFile := filepath.Join("{{path}}", "{{table}}.go")
			handleFile = tools.ReplaceAllData(handleFile, map[string]string{
				"{{path}}":  handlePath,
				"{{table}}": table,
			})
			//判断文件存在，如果存在 就备份之前文件
			if tools.CheckFileIsExist(handleFile) {
				tools.MustCheck(os.Rename(handleFile, fmt.Sprintf("%s.bak", handleFile)))
			}
			//向文件中写入数据
			tools.WriteToFile(handleFile, handleText)
			xconsole.Greenf("handle", handleFile)
		}

		//server
		{
			var serverText string
			//获取到模板文件
			serverText, err = tools.ParseTmplFile(filepath.Join(dir, "tmp", options.ProjectName, "model", "server.go.tmpl"), map[string]string{
				"ProjectName": options.ProjectName,
				"appName":     options.AppName,
				"AppName":     tools.StrFirstToUpper(options.AppName),
				"TableName":   modelName,
				"tableName":   table,
				"Id":          Id,
				"id":          tools.UnStrFirstToUpper(Id),
			})
			tools.MustCheck(err)
			//模板替换文件位置
			//模板替换文件夹位置
			serverPath := filepath.Join(dir, "internal", "{{appName}}", "services", "{{table}}")
			serverPath = tools.ReplaceAllData(serverPath, map[string]string{
				"{{appName}}": options.AppName,
				"{{table}}":   table,
			})
			//创建文件夹
			tools.MustCheck(os.MkdirAll(serverPath, 777))
			//模板替换文件位置
			serverFile := filepath.Join("{{path}}", "{{table}}.go")
			serverFile = tools.ReplaceAllData(serverFile, map[string]string{
				"{{path}}":  serverPath,
				"{{table}}": table,
			})
			//判断文件存在，如果存在 就备份之前文件
			if tools.CheckFileIsExist(serverFile) {
				tools.MustCheck(os.Rename(serverFile, fmt.Sprintf("%s.bak", serverFile)))
			}
			//向文件中写入数据
			tools.WriteToFile(serverFile, serverText)
			xconsole.Greenf("services", serverFile)
		}

		//registry
		{
			var registryText string
			//获取到模板文件
			registryText, err = tools.ParseTmplFile(filepath.Join(dir, "tmp", options.ProjectName, "model", "registry.go.tmpl"), map[string]string{
				"ProjectName": options.ProjectName,
				"appName":     options.AppName,
				"TableName":   modelName,
				"tableName":   table,
				"version":     tools.UnStrFirstToUpper(options.Version),
			})
			tools.MustCheck(err)
			//模板替换文件位置
			registryPath := filepath.Join(dir, "internal", "{{appName}}", "api", "{{version}}", "registry")
			registryPath = tools.ReplaceAllData(registryPath, map[string]string{
				"{{appName}}": options.AppName,
				"{{version}}": tools.UnStrFirstToUpper(options.Version),
			})
			//创建文件夹
			tools.MustCheck(os.MkdirAll(registryPath, 777))
			registryFile := filepath.Join("{{path}}", "{{table}}.go")
			registryFile = tools.ReplaceAllData(registryFile, map[string]string{
				"{{path}}":  registryPath,
				"{{table}}": table,
			})
			//判断文件存在，如果存在 就备份之前文件
			if tools.CheckFileIsExist(registryFile) {
				tools.MustCheck(os.Rename(registryFile, fmt.Sprintf("%s.bak", registryFile)))
			}
			//向文件中写入数据
			tools.WriteToFile(registryFile, registryText)
			xconsole.Greenf("registry", registryFile)
		}

		//map
		{
			var mapText string
			//获取到模板文件
			mapText, err = tools.ParseTmplFile(filepath.Join(dir, "tmp", options.ProjectName, "model", "map.go.tmpl"), map[string]string{
				"TableName":     modelName,
				"TableFieldMap": TableFieldMap,
			})
			tools.MustCheck(err)
			//模板替换文件位置
			mapFile := filepath.Join(dir, "internal", "{{appName}}", "map", "{{table}}.go")
			mapFile = tools.ReplaceAllData(mapFile, map[string]string{
				"{{appName}}": options.AppName,
				"{{table}}":   table,
			})
			//判断文件存在，如果存在 就备份之前文件
			if tools.CheckFileIsExist(mapFile) {
				tools.MustCheck(os.Rename(mapFile, fmt.Sprintf("%s.bak", mapFile)))
			}
			//向文件中写入数据
			tools.WriteToFile(mapFile, mapText)
			xconsole.Greenf("map", mapFile)
		}

		//关闭sql链接
		listRows.Close()
	}
}
