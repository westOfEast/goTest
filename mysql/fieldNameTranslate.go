package mysql

import (
	"strings"
)

func FieldNameTranslate(fieldName string) string {
	result := fieldName
	if "" != fieldName {
		tmpResult := ""
		splitResult := strings.Split(fieldName, "_")
		if len(splitResult) > 1 {
			for index := 0; index < len(splitResult); index++ {
				if index > 0 {
					tmpResult += upperFirstLetter(splitResult[index])
				} else {
					tmpResult += splitResult[index]
				}
			}
		}
		result = tmpResult
	}
	return result
}

func upperFirstLetter(source string) string {
	runeArray := []rune(source)
	result := ""
	for runeIndex := 0; runeIndex < len(runeArray); runeIndex++ {
		if runeIndex == 0 {
			runeArray[runeIndex] = runeArray[runeIndex] - 32
		}
		result += string(runeArray[runeIndex])
	}
	if "" == result {
		result = source
	}
	return result
}
