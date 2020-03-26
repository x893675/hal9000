package database

import (
	"github.com/fatih/structs"
	"hal9000/pkg/utils/stringutil"
)

func GetColumnsFromStruct(s interface{}) []string {
	names := structs.Names(s)
	for i, name := range names {
		names[i] = stringutil.CamelCaseToUnderscore(name)
	}
	return names
}

func GetColumnsFromStructWithPrefix(prefix string, s interface{}) []string {
	names := structs.Names(s)
	for i, name := range names {
		names[i] = WithPrefix(prefix, stringutil.CamelCaseToUnderscore(name))
	}
	return names
}

func WithPrefix(prefix, str string) string {
	return prefix + "." + str
}