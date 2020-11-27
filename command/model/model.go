package model

import (
	"fmt"
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

var modelTpl = `
package app

import (
	"{{ProjectName}}/internal/{{AppName}}/model"
	"github.com/jinzhu/gorm"
	{{IsTime}}
)

func init() {
	model.MainDB.AutoMigrate(new({{TableName}}))
}

type {{TableName}} struct {
	{{TableFieldList}}
}

func (a *{{TableName}}) TableName() string {
	return "{{tableName}}"
}

//添加
func (a *{{TableName}}) Add() error {
	return model.MainDB.Table(a.TableName()).Create(a).Error
}

//删除where
func (a *{{TableName}}) Del(wheres map[string]interface{}) error {
	db := model.MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	return db.Delete(a).Error
}

//查询所有
func (a *{{TableName}}) GetAll() (data []App, err error) {
	err = model.MainDB.Table(a.TableName()).Find(&data).Error
	return
}

//偏移查询
func (a *{{TableName}}) Get(start int64, size int64, wheres map[string]interface{}) (data []App, total int64, err error) {
	db := model.MainDB.Table(a.TableName())
	for k, v := range wheres {
		db = db.Where(k, v)
	}
	err = db.Limit(size).Offset(start).Find(&data).Error
	err = db.Count(&total).Error
	return
}

//根据id查询
func (a *{{TableName}}) GetById() error {
	return model.MainDB.Table(a.TableName()).Where("id=?", a.ID).First(a).Error
}

//修改ById
func (a *{{TableName}}) UpdateById() error {
	return model.MainDB.Table(a.TableName()).Where("id=?", a.ID).Update(a).Error
}
`

func (options *RunOptions) Run() {
	var err error
	cfg, err := config.TryLoadFromDisk()
	tools.MustCheck(err)
	db, err := database.NewDatabaseClient(cfg.Mysql, nil)
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
		text = tools.ReplaceAllData(modelTpl, map[string]string{
			"{{TableFieldList}}": TableFieldList,
			"{{ProjectName}}":    options.ProjectName,
			"{{AppName}}":        options.AppName,
			"{{TableName}}":      modelName,
			"{{tableName}}":      table,
		})
		fmt.Println(isTime)
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
