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
	cfg, err := config.TryLoadFromDisk()
	tools.MustCheck(err)
	db, err := database.NewDatabaseClient(cfg.Mysql, nil)
	tools.MustCheck(err)

	//git
	tools.MustCheck(tools.GitClone("https://github.com/myxy99/Yangon-tpl.git", "tmp\\"+options.ProjectName))
	modelTpl, err := ioutil.ReadFile(fmt.Sprintf(`tmp/%s/model/model.go`, options.ProjectName))
	defer tools.RemoveAllList("tmp")
	tools.MustCheck(err)
	rows, err := db.DB().Raw("show tables;").Rows()
	tools.MustCheck(err)
	defer rows.Close()
	var table string
	for rows.Next() {
		tools.MustCheck(rows.Scan(&table))
		modelName := tools.StrFirstToUpper(tools.Capitalize(table))
		listRows, err := db.DB().Raw(fmt.Sprintf("show columns from %s;", table)).Rows()
		tools.MustCheck(err)
		var TableFieldList, text string
		isTime := false
		for listRows.Next() {
			list := new(List)
			_ = listRows.Scan(&list.Key, &list.Type, &list.Default, &list.Extra, &list.Field, &list.Null)
			//所有的字段
			var structType string
			var tmpIsTime bool
			structType, tmpIsTime = tools.SqlType2StructType(list.Type)
			isTime = tmpIsTime || isTime
			TableFieldList += fmt.Sprintf("%s\t%s\n\t", tools.StrFirstToUpper(tools.Capitalize(list.Key)), structType)
		}
		text = tools.ReplaceAllData(string(modelTpl), map[string]string{
			"{{TableFieldList}}": TableFieldList,
			"{{ProjectName}}":    options.ProjectName,
			"{{AppName}}":        options.AppName,
			"{{TableName}}":      modelName,
			"{{tableName}}":      table,
		})
		if isTime {
			text = strings.ReplaceAll(text, "{{IsTime}}", "\"time\"")
		} else {
			text = strings.ReplaceAll(text, "{{IsTime}}", "")
		}
		path := `internal/{{AppName}}/model/{{table}}`
		path = tools.ReplaceAllData(path, map[string]string{
			"{{AppName}}": options.AppName,
			"{{table}}":   table,
		})
		_ = os.MkdirAll(path, 777)
		file := `{{path}}/{{table}}.go`
		file = tools.ReplaceAllData(file, map[string]string{
			"{{path}}":  path,
			"{{table}}": table,
		})
		tools.WriteToFile(file, text)
		fmt.Println(file)
	}
}
