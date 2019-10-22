package main

import (
	"log"
	"os"
	"rds_backup/util"
)

/**
@auther funnyzpc
@description 命令行方式下载文件(需打包,使用命令行方式执行)
@cmd > 可执行文件 下载地址 文件名称(含目录地址)
*/
func main() {
	if len(os.Args) != 3 {
		log.Println("命令行参数有误!")
		os.Exit(1)
	}
	url := os.Args[1]
	filename := os.Args[2]
	err := util.DownloadFile(url, filename)
	if err != nil {
		log.Println("文件下载异常:", err)
		return
	}
	log.Println("文件下载成功!")
}
