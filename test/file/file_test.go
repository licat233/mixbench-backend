package file_test

import (
	"mixbench/common/utils"
	"path"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	root := utils.CurrentProjectRootPath()
	filename := path.Join(root, "deploy/email/template/captcha.html")
	println("路径:", filename)
}
