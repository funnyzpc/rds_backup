package service

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"rds_backup/util"
	"time"
)

func LogBakup01() bool {
	log.Println("*****开始处理日志备份01******")

	// 构建配置
	config := &ssh.ClientConfig{
		User:            "SSH用户名",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("SSH密码"),
		},
	}
	// 构建TCP连接
	client, err := ssh.Dial("tcp", "服务器地址:端口", config)
	if err != nil {
		log.Println("TCP 连接失败: " + err.Error())
		return false
	}
	log.Println("连接成功")

	// 构建sftp连接
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer sftp.Close()

	// 构建目录
	time := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	targetPath := "D:/logs_bak/product_bak/" + time
	if false == util.DirExists(targetPath) {
		log.Println("构建目录失败01")
		return false
	}

	// 下载备份文件
	process011(targetPath, time, sftp)
	process012(targetPath, time, sftp)
	process013(targetPath, time, sftp)
	process014(targetPath, time, sftp)

	// 循环目录并解压文件(略)
	log.Println("*****日志备份完成01******")
	return true
}

func process011(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 01 *********************************
	srcFile, err := sftp.Open("/路径/文件01.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败01:", err)
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(targetPath + "/文件01.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败01:", err)
		return
	}
	defer dstFile.Close()
	srcFile.WriteTo(dstFile)
}

func process012(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 02 *********************************
	srcFile02, err := sftp.Open("/路径/文件02.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败02:", err)
		return
	}
	defer srcFile02.Close()
	dstFile02, err := os.Create(targetPath + "/文件02.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败02:", err)
		return
	}
	defer dstFile02.Close()
	srcFile02.WriteTo(dstFile02)
}

func process013(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 03 *********************************
	srcFile03, err := sftp.Open("/路径/文件03.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败03:", err)
		return
	}
	defer srcFile03.Close()
	dstFile03, err := os.Create(targetPath + "/文件03.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败03:", err)
		return
	}
	defer dstFile03.Close()
	srcFile03.WriteTo(dstFile03)
}

func process014(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 04 *********************************
	srcFile04, err := sftp.Open("/路径/文件04.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败04:", err)
		return
	}
	defer srcFile04.Close()
	dstFile04, err := os.Create(targetPath + "/文件04.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败04:", err)
		return
	}
	defer dstFile04.Close()
	srcFile04.WriteTo(dstFile04)
}
