package util

import (
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

/**
	@description 解压缩文件
	@fullZipFile zip文件所在的全路径
	@releaseDir 解压文件释放路径
	@password 解压密码
**/
func UnzipWithPasword(fullZipFile string, releaseDir string, password string) bool {
	r, err := zip.OpenReader(fullZipFile)
	if err != nil {
		log.Println("ZIP文件读取出错:", err)
		return false
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}
		writeFile(releaseDir, f)
	}
	return true
}

func writeFile(filePath string, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err := rc.Close(); err != nil {
			panic(err)
		}
	}()
	path := filepath.Join(filePath, f.Name)
	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
	} else {
		os.MkdirAll(filepath.Dir(path), f.Mode())
		fo, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(fo, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
