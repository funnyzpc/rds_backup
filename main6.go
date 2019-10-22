package main

import (
	"log"
	"rds_backup/util"
)

/**
@auther funnyzpc
@description 测试带密码的zip文件解压
*/
func main() {
	filePath := "C:\\Users\\Administrator\\Desktop\\Desktop.zip"
	releaseDir := "C:\\Users\\Administrator\\Desktop\\tt"
	util.UnzipWithPasword(filePath, releaseDir, "123qwe")
	log.Println("=====>解压成功!")
}
