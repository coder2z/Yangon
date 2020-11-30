package model

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"yangon/command/model/config"
	"yangon/pkg/database"
	"yangon/tools"
)

type List struct {
	Field   string `gorm:"Field"`
	Type    string `gorm:"Type"`
	Null    string `gorm:"Null"`
	Key     string `gorm:"Key"`
	Default string `gorm:"Default"`
	Extra   string `gorm:"Extra"`
}

func (options *RunOptions) Run() {
	var err error
	//解析配置
	cfg, err := config.TryLoadFromDisk()
	tools.MustCheck(err)
	//链接数据库
	db, err := database.NewDatabaseClient(cfg.Mysql, nil)
	tools.MustCheck(err)
	//拉取模板
	tools.MustCheck(tools.GitClone("https://github.com/myxy99/Yangon-tpl.git", "tmp\\"+options.ProjectName))
	//获取到模板文件
	modelTpl, err := ioutil.ReadFile(fmt.Sprintf(`tmp/%s/model/model.go`, options.ProjectName))
	//defer删除拉取的模板
	defer tools.RemoveAllList("tmp")
	tools.MustCheck(err)
	//查找表
	rows, err := db.DB().Raw("show tables;").Rows()
	tools.MustCheck(err)
	defer rows.Close()
	var table string
	for rows.Next() {
		tools.MustCheck(rows.Scan(&table))
		//把表名进行驼峰式转换
		modelName := tools.StrFirstToUpper(tools.Capitalize(table))
		//查字段名
		listRows, err := db.DB().Raw(fmt.Sprintf("show columns from %s;", table)).Rows()
		tools.MustCheck(err)
		var TableFieldList, text string
		isTime := false
		Id := "Id"
		for listRows.Next() {
			list := new(List)
			//查询所有字段
			_ = listRows.Scan(&list.Key, &list.Type, &list.Default, &list.Extra, &list.Field, &list.Null)
			//把字段名进行驼峰式转换
			upper := tools.StrFirstToUpper(tools.Capitalize(list.Key))
			//判断是不是主键
			if tools.IsPRI(list.Type) {
				Id = upper
			}
			var structType string
			var tmpIsTime bool
			//把对应的字段类型转换为结构体类型
			structType, tmpIsTime = tools.SqlType2StructType(list.Type)
			isTime = tmpIsTime || isTime
			//组合结构体中的字段，字符串
			TableFieldList += fmt.Sprintf("%s\t%s\n\t", upper, structType)
		}
		//模板替换
		text = tools.ReplaceAllData(string(modelTpl), map[string]string{
			"{{TableFieldList}}": TableFieldList,
			"{{ProjectName}}":    options.ProjectName,
			"{{AppName}}":        options.AppName,
			"{{TableName}}":      modelName,
			"{{tableName}}":      table,
			"{{ID}}":             Id,
		})
		//是否使用了time 包
		if isTime {
			text = strings.ReplaceAll(text, "{{IsTime}}", "\"time\"")
		} else {
			text = strings.ReplaceAll(text, "{{IsTime}}", "")
		}
		//模板替换文件夹位置
		path := `internal/{{AppName}}/model/{{table}}`
		path = tools.ReplaceAllData(path, map[string]string{
			"{{AppName}}": options.AppName,
			"{{table}}":   table,
		})
		//创建文件夹
		tools.MustCheck(os.MkdirAll(path, 777))
		//模板替换文件位置
		file := `{{path}}/{{table}}.go`
		file = tools.ReplaceAllData(file, map[string]string{
			"{{path}}":  path,
			"{{table}}": table,
		})
		//判断文件存在，如果存在 就备份之前文件
		if tools.CheckFileIsExist(file) {
			tools.MustCheck(os.Rename(file, fmt.Sprintf("%s.bak", file)))
		}
		//向文件中写入数据
		tools.WriteToFile(file, text)
		fmt.Println(file)
		//关闭sql链接
		listRows.Close()
	}
}
