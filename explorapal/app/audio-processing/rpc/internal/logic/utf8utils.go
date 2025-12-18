package logic

import (
	"strings"
	"unicode/utf8"
)

// sanitizeUTF8 清理字符串中的无效UTF-8字符
func sanitizeUTF8(s string) string {
	if utf8.ValidString(s) {
		return s
	}
	// 将无效UTF-8字符替换为有效的Unicode替换字符
	return strings.ToValidUTF8(s, "�")
}

// sanitizeUTF8Slice 清理字符串切片中的无效UTF-8字符
func sanitizeUTF8Slice(slice []string) []string {
	if slice == nil {
		return nil
	}
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = sanitizeUTF8(s)
	}
	return result
}
