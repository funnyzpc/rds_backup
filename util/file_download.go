package util

import (
	"io"
	"log"
	"net/http"
	"os"
)

/**
url: 下载地址
filePath: 文件目录
*/
func DownloadFile(url string, filePath string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("错误: ", err)
		return err
	}
	return nil
}
