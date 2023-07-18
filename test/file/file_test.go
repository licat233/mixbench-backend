package file_test

import (
	"fmt"
	"mixbench/common/utils"
	"path"
	"regexp"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	root := utils.ProjectRoot
	filename := path.Join(root, "deploy/email/template/captcha.html")
	println("路径:", filename)
}

func TestGetprojectroot(t *testing.T) {
	// ex := utils.CurrentPath()

	// fmt.Println(ex)

	fullPath := "/Users/licat/Desktop/mixbench/mixbench-backend/common/emailx/xxxx"

	reg := regexp.MustCompile("^.*mixbench-backend")
	matched := reg.FindString(fullPath)
	// rootDir := strings.TrimSuffix(matched, "common")

	fmt.Println(matched)
}
