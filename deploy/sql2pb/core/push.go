/*
 * @Author: licat
 * @Date: 2023-01-30 09:43:14
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 14:43:58
 * @Description: licat233@gmail.com
 */
package core

import (
	"fmt"
	"os"
	"strings"
)

const (
	customMessageStartMark = "//------ custom message start ------"
	customMessageEndMark   = "//------ custom message end ------"

	customServiceStartMark = "//------ custom service start ------"
	customServiceEndMark   = "//------ custom service end ------"
)

/**
 * @description: 生成并更新文件
 * @param {string} filename
 * @param {string} content
 * @return {error}
 */
func Push(filename, content string) error {
	//read
	fd, _ := os.ReadFile(filename)

	//open or create and replace
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer f.Close()

	//find
	fileContent := string(fd)

	//update message
	content, err = updateCustomMessage(fileContent, content)
	if err != nil {
		return err
	}
	//update service
	content, err = updateCustomService(fileContent, content)
	if err != nil {
		return err
	}

	//write
	_, err = f.WriteString(content)
	return err
}

/**
 * @description:更新 custom Message
 * @param {string} oldContent
 * @param {string} newContent
 * @return {string, error}
 */
func updateCustomMessage(oldContent, newContent string) (string, error) {
	//statrt提取custom message
	startIndex := strings.Index(oldContent, customMessageStartMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customMessageStartMark)
	}
	endIndex := strings.Index(oldContent, customMessageEndMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customMessageEndMark)
	}
	customMessage := oldContent[startIndex+len(customMessageStartMark) : endIndex]
	//end提取custom message

	//start分割节点
	startIndex = strings.Index(newContent, customMessageStartMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customMessageStartMark)
	}
	endIndex = strings.Index(newContent, customMessageEndMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customMessageEndMark)
	}
	topContent := newContent[:startIndex+len(customMessageStartMark)]
	bottomContent := newContent[endIndex:]
	//end分割节点

	return fmt.Sprintf("%s%s%s", topContent, customMessage, bottomContent), nil
}

/**
 * @description: 更新 custom Service
 * @param {string} oldContent
 * @param {string} newContent
 * @return {string, error}
 */
func updateCustomService(oldContent, newContent string) (string, error) {
	startIndex := strings.Index(oldContent, customServiceStartMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customServiceStartMark)
	}
	endIndex := strings.Index(oldContent, customServiceEndMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customServiceEndMark)
	}

	customService := oldContent[startIndex+len(customServiceStartMark) : endIndex]

	startIndex = strings.Index(newContent, customServiceStartMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customServiceStartMark)
	}
	endIndex = strings.Index(newContent, customServiceEndMark)
	if startIndex == -1 {
		return errorHandler(oldContent, newContent, customServiceEndMark)
	}

	topContent := newContent[:startIndex+len(customServiceStartMark)]
	bottomContent := newContent[endIndex:]

	return fmt.Sprintf("%s%s%s", topContent, customService, bottomContent), nil
}

func errorHandler(oldContent, newContent, mark string) (string, error) {
	if strings.TrimSpace(oldContent) == "" {
		return newContent, nil
	}
	return "", fmt.Errorf("not found %s", mark)
}
