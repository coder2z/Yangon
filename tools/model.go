package tools

import (
	"fmt"
	"regexp"
	"strings"
)

var EImportsHead = map[string]string{
	"time.Time":      `"time"`,
	"datatypes.JSON": `"gorm.io/datatypes"`,
	"datatypes.Date": `"gorm.io/datatypes"`,
}

func SqlType2StructType(t, isnull string) (string, bool) {
	isnullB := strings.EqualFold(isnull, "YES")
	if v, ok := TypeMysqlDicMp[t]; ok {
		return fixNullToPorint(v, isnullB), regexp.MustCompile(`date`).MatchString(t)
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for _, l := range TypeMysqlMatchList {
		if ok, _ := regexp.MatchString(l.Key, t); ok {
			return fixNullToPorint(l.Value, isnullB), regexp.MustCompile(`date`).MatchString(t)
		}
	}
	panic(fmt.Sprintf("type (%v) not match in any way)", t))
}

func fixNullToPorint(name string, isNull bool) string {
	if isNull {
		if strings.HasPrefix(name, "uint") {
			return "*" + name
		}
		if strings.HasPrefix(name, "int") {
			return "*" + name
		}
		if strings.HasPrefix(name, "float") {
			return "*" + name
		}
	}
	return name
}

func IsPRI(t string) bool {
	return regexp.MustCompile(`PRI`).MatchString(t)
}

// TypeMysqlDicMp Accurate matching type.精确匹配类型
var TypeMysqlDicMp = map[string]string{
	"smallint":            "int16",
	"smallint unsigned":   "uint16",
	"int":                 "int",
	"int unsigned":        "uint",
	"bigint":              "int64",
	"bigint unsigned":     "uint64",
	"varchar":             "string",
	"char":                "string",
	"date":                "datatypes.Date",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int8",
	"tinyint unsigned":    "uint8",
	"tinyint(1)":          "bool", // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "bool", // tinyint(1) 默认设置成bool
	"json":                "datatypes.JSON",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
}

// TypeMysqlMatchList Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchList = []struct {
	Key   string
	Value string
}{
	{`^(tinyint)[(]\d+[)] unsigned`, "uint8"},
	{`^(smallint)[(]\d+[)] unsigned`, "uint16"},
	{`^(int)[(]\d+[)] unsigned`, "uint32"},
	{`^(bigint)[(]\d+[)] unsigned`, "uint64"},
	{`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(tinyint)[(]\d+[)]`, "int8"},
	{`^(smallint)[(]\d+[)]`, "int16"},
	{`^(int)[(]\d+[)]`, "int"},
	{`^(bigint)[(]\d+[)]`, "int64"},
	{`^(char)[(]\d+[)]`, "string"},
	{`^(enum)[(](.)+[)]`, "string"},
	{`^(varchar)[(]\d+[)]`, "string"},
	{`^(varbinary)[(]\d+[)]`, "[]byte"},
	{`^(blob)[(]\d+[)]`, "[]byte"},
	{`^(binary)[(]\d+[)]`, "[]byte"},
	{`^(decimal)[(]\d+,\d+[)]`, "float64"},
	{`^(mediumint)[(]\d+[)]`, "string"},
	{`^(double)[(]\d+,\d+[)]`, "float64"},
	{`^(float)[(]\d+,\d+[)]`, "float64"},
	{`^(datetime)[(]\d+[)]`, "time.Time"},
	{`^(bit)[(]\d+[)]`, "[]uint8"},
	{`^(text)[(]\d+[)]`, "string"},
	{`^(integer)[(]\d+[)]`, "int"},
	{`^(timestamp)[(]\d+[)]`, "time.Time"},
}
