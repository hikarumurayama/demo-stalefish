package char_filter

import "strings"

type CharFilter interface {
	Filter(string) string
}

type MappingCharFilter struct {
	Mapper map[string]string
}

// あらかじめマッピングされた文字列に変換する
func (c MappingCharFilter) Filter(s string) string {
	for k, v := range c.Mapper {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
