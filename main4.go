package main

import (
	"log"
	"rds_backup/service"
	"time"
)

/**
	@auther funnyzpc
	@description 备份日志文件
**/
func main() {
	time := time.Now().Format("2006-01-02 15:04")

	log.Println(">>>>>>>>>>", time, "日志备份开始<<<<<<<<<<")

	if !service.LogBakup01() {
		log.Fatal("日志01备份异常~")
	}

	if !service.LogBakup02() {
		log.Fatal("日志02备份异常~")
	}
	log.Println(">>>>>>>>>>", time, "日志备份结束<<<<<<<<<<")
}
