package utils

import (
	"os"
	"strings"
)

func NewStringByEnv(str string) string {
	// 获取字符串中的环境变量
	envVariables := getEnvironmentVariables(str)

	// 替换字符串中的环境变量
	result := replaceEnvironmentVariables(str, envVariables)
	return result
}

// 获取字符串中的环境变量
func getEnvironmentVariables(str string) []string {
	var envVariables []string

	for _, match := range extractEnvironmentVariables(str) {
		// 提取环境变量名称
		envVar := match[2 : len(match)-1]

		// 检查环境变量是否存在
		if value := os.Getenv(envVar); value != "" {
			envVariables = append(envVariables, match)
		}
	}

	return envVariables
}

// 替换字符串中的环境变量
func replaceEnvironmentVariables(str string, envVariables []string) string {
	for _, match := range envVariables {
		// 提取环境变量名称
		envVar := match[2 : len(match)-1]

		// 获取环境变量的值
		value := os.Getenv(envVar)

		// 替换字符串中的环境变量值
		str = strings.Replace(str, match, value, -1)
	}

	return str
}

// 提取字符串中的环境变量
func extractEnvironmentVariables(str string) []string {
	var envVariables []string

	startIndex := -1
	for i := 0; i < len(str); i++ {
		// 检查是否匹配环境变量的起始字符"${"
		if str[i] == '$' && i+1 < len(str) && str[i+1] == '{' {
			startIndex = i
		}

		// 检查是否匹配环境变量的结束字符"}"
		if startIndex != -1 && str[i] == '}' {
			envVariables = append(envVariables, str[startIndex:i+1])
			startIndex = -1
		}
	}

	return envVariables
}
