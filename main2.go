package main

import (
	"log"
	"rds_backup/util"
	"time"
)

/**
@auther funnyzpc
@description 测试tar.gz文件解压缩
*/
func main() {
	gzFilePath := "C:\\Users\\Administrator\\Desktop\\testzip.tar.gz"
	suffix := time.Now().Format("2006-01-02")
	unGzFilePath := "D:\\tmp\\" + suffix
	err := util.UnTarGz(gzFilePath, unGzFilePath)
	if nil != err {
		log.Println("解压错误:", err)
		return
	}
	log.Println("解压成功!")
}
