package util

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
)

/**
srcFilePath: 压缩包路径
destDirPath： 解压文件释放目录
*/
func UnTarGz(srcFilePath string, destDirPath string) error {
	log.Println("UnTarGz Params  :", srcFilePath, destDirPath)
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if nil != err {
		log.Println("打开文件失败 %s", err)
		return err
	}
	defer fr.Close()
	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if nil != err {
		log.Println("读取文件失败 %s", err)
		// return err
	}
	// Tar reader
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if hdr == nil {
			log.Println("结束读取文件!")
			break
		}
		if err == io.EOF {
			log.Println("读取文件失败: %s", err)
			break
			// return err
		}
		//handleError(err)
		log.Println("UnTarGz file: " + hdr.Name)
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to file
			fw, err := os.Create(destDirPath + "/" + hdr.Name)
			if nil != err {
				log.Println("创建异常:", err)
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				log.Println("拷贝异常:", err)
			}
		}
	}
	log.Println("解压完成 !")
	return nil
}
