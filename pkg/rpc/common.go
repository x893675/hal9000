package rpc

import (
	"github.com/fatih/structs"
	"strings"
)

const (
	TagName              = "json"
	SearchWordColumnName = "search_word"
)

type Request interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)
}

func getFieldName(field *structs.Field) string {
	tag := field.Tag(TagName)
	t := strings.Split(tag, ",")
	if len(t) == 0 {
		return "-"
	}
	return t[0]
}