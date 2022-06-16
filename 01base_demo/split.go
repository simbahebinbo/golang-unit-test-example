package base_demo

import "strings"

// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
func Split(s, sep string) (result []string) {
	//返回子串str在字符串s中第一次出现的位置。
	//如果找不到则返回-1；如果str为空，则返回0
	i := strings.Index(s, sep)
	result = make([]string, 0, strings.Count(s, sep)+1)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
