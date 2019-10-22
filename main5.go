package main

import (
	"log"
	"rds_backup/util"
)

/**
@auther funnyzpc
@description 测试zip文件解压
*/
func main() {
	fullZipFile := "http://mirror.rise.ph/apache/tomcat/tomcat-9/v9.0.27/bin/apache-tomcat-9.0.27.zip"
	releaseDir := "D:/tmp/download"
	err := util.Unzip(fullZipFile, releaseDir)
	if err != nil {
		log.Println("zip文件解压失败:", err)
		return
	}
	log.Println("zip文件解压成功")
}
