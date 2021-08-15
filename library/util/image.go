package util

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"fast-gin/config"
	"fast-gin/library/file"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return config.Cfg.App.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return config.Cfg.App.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return config.Cfg.App.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range config.Cfg.App.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println("CheckImageSize err====", err)
		return false
	}
	return size <= config.Cfg.App.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
