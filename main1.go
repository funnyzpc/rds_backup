package main

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"log"
	"rds_backup/util"
	"time"
)

/**
@auther funnyzpc
@description 定时任务拉取RDS备份文件
*/
func main() {
	log.Println("******开始备份DB******")
	// nowTime := time.Now().Format("2006-01-02 15:04:05")
	client, err := rds.NewClientWithAccessKey("cn-hangzhou", "您的accessKeyId", "您的accessKeySecret")
	if err != nil {
		panic(err)
	}

	request := rds.CreateDescribeBackupsRequest()
	request.StartTime = time.Now().AddDate(0, 0, -1).UTC().Format(time.RFC3339)[0:11] + "00:00Z"
	request.DBInstanceId = "您的实例ID"
	request.PageSize = "30"
	request.PageNumber = "1"
	request.BackupStatus = "Success"
	request.Scheme = "https"

	respJsonStr, err := client.DescribeBackups(request)
	log.Println("api响应信息：", respJsonStr.GetHttpContentString())
	if err != nil {
		log.Println("api获取备份文件错误:", err.Error())
		return
	}
	data := &BackupInfo{}
	if json.Unmarshal(respJsonStr.GetHttpContentBytes(), data) != nil {
		log.Println("json to struct error !")
		return
	}
	if data.TotalRecordCount > 0 {
		// 获取下载地址
		downloadUrl := data.Items.BackupItem[0].BackupDownloadURL
		backupDate := data.Items.BackupItem[0].BackupStartTime[0:10]
		log.Println("DB备份时间: ", backupDate)

		log.Println("DB备份 download url: ", downloadUrl)
		// 下载文件
		// suffix := time.Now().Format("2006-01-02")
		filePath := "D:\\db_bak\\product\\" + "product_all_" + backupDate + ".tar.gz"
		err := util.DownloadFile(downloadUrl, filePath)
		if err != nil {
			log.Println("下载失败:", err)
			return
		}
		//解压全量备份文件
		unTarGzFilePath := "D:\\db_bak\\product\\" + backupDate
		err2 := util.UnTarGz(filePath, unTarGzFilePath)
		if err2 != nil {
			log.Println("解压文件失败: ", err2)
			return
		}
		//执行window CMD命令 导入DB FILE(略)
		/*
			f, err := exec.Command("ls", "/").Output()
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(string(f))
		*/
		log.Println("******获取、下载、解压 文件成功******")
	} else {
		log.Println("备份记录为空,停止操作!")
	}
}

type BackupInfo struct {
	Items            *BackupItems `json:"items"`
	TotalRecordCount int          `json:"totalRecordCount"`
}

type BackupItems struct {
	BackupItem *[2]Backup `json:"backup"`
}

type Backup struct {
	BackupStartTime           string
	BackupEndTime             string
	BackupSize                int32
	BackupStatus              string `json:"backupStatus"`
	BackupIntranetDownloadURL string `json:"backupIntranetDownloadURL"`
	BackupDownloadURL         string `json:"backupDownloadURL"`
}
