package service

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"rds_backup/util"
	"time"
)

func LogBakup02() bool {
	log.Println("*****开始处理日志备份02******")

	// 构建配置
	config := &ssh.ClientConfig{
		User:            "SSH用户名",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password("SSH密码"),
		},
	}
	// 构建TCP连接
	client, err := ssh.Dial("tcp", "服务器地址:服务器IP", config)
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
		return false
	}

	// 下载单个文件或zip
	process01(targetPath, time, sftp)
	process02(targetPath, time, sftp)
	process03(targetPath, time, sftp)
	process04(targetPath, time, sftp)
	process05(targetPath, time, sftp)
	process06(targetPath, time, sftp)

	return true
}

func process01(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 01 *********************************
	srcFile, err := sftp.Open("/路径/" + time + "/文件01.log." + time + ".zip")
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

func process02(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 02 *********************************
	srcFile02, err := sftp.Open("/路径/" + time + "/文件02.log." + time + ".zip")
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

func process03(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 03 *********************************
	srcFile03, err := sftp.Open("/路径/" + time + "/文件03.log." + time + ".zip")
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

func process04(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 04 *********************************
	srcFile04, err := sftp.Open("/路径/" + time + "/文件04.log." + time + ".zip")
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

func process05(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 05 *********************************
	srcFile05, err := sftp.Open("/路径/" + time + "/文件05.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败05:", err)
		return
	}
	defer srcFile05.Close()
	dstFile05, err := os.Create(targetPath + "/文件05.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败05:", err)
		return
	}
	defer dstFile05.Close()
	srcFile05.WriteTo(dstFile05)
}

func process06(targetPath string, time string, sftp *sftp.Client) {
	// sftp传输文件 06 *********************************
	srcFile06, err := sftp.Open("/路径/" + time + "/文件06.log." + time + ".zip")
	if err != nil {
		log.Println("文件打开失败06:", err)
		return
	}
	defer srcFile06.Close()
	dstFile06, err := os.Create(targetPath + "/文件06.log." + time + ".zip")
	if err != nil {
		log.Println("文件传输失败06:", err)
		return
	}
	defer dstFile06.Close()
	srcFile06.WriteTo(dstFile06)

	// 循环目录并解压文件(略)
	log.Println("*****日志备份完成02******")
}
