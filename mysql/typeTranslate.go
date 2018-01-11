package mysql

var typeMap = map[string]string{
	"varchar":  "string",
	"double":   "Double",
	"bigint":   "Long",
	"int":      "Integer",
	"datetime": "Date",
	"tinyint":  "Integer",
	"date":     "Date",
	"float":    "Double",
}

func TypeTranslate(fieldType string) string {
	value, ok := typeMap[fieldType]
	if !ok {
		value = "String"
	}
	return value
}
